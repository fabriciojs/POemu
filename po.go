package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"encoding/json"
	"time"
	"strings"
	"strconv"
	"flag"
)

var searchjson = flag.String("searchjson", "", "the complete search JSON")

var media = flag.String("media", "", "Media Code [1, 14]")
var query = flag.String("query", "", "Keyword to query")
var limit = flag.String("limit", "2000", "limit items amount, first import")
var search_port = flag.String("rpc_port", "8222", "The RPC port to connect to the search service")

var cod_busca = flag.String("cod_busca", "", "o Código da busca")
var cod_monitoramento = flag.String("cod_monitoramento", "0", "o Código do monitoramento")
var cod_campanha = flag.String("cod_campanha", "0", "o Código do campanha")
var cod_conta = flag.String("cod_conta", "", "o Código da conta")

func main() {
	flag.Parse()

	search := make(map[string]interface{})

	if *searchjson == "" {
		if *query == "" {
			fmt.Println("You must provide the --query param.")
			os.Exit(1)
		}

		if *cod_conta == "" {
			fmt.Println("You must provide the --cod_conta param.")
			os.Exit(1)
		}

		if *cod_monitoramento == "0" && *cod_campanha == "0" {
			fmt.Println("You must provide the --cod_monitoramento param.")
			os.Exit(1)
		}

		if *cod_busca == "" {
			fmt.Println("You must provide the --cod_busca param.")
			os.Exit(1)
		}

		if *media == "" {
			fmt.Println("You must provide the --media param.")
			os.Exit(1)
		}

		search["LastItem"] = ""
		search["FormattedResult"] = "0"
		search["First"] = "1"
		search["Feed"] = "http://livebuzz.com.br/dev/importapi/search/store"
		search["Extid"] = fmt.Sprintf("%s.%s.%s.%s.%s", *cod_conta, *cod_monitoramento, *cod_campanha, *cod_busca, *media)
		search["Limit"] = *limit
		search["Media"] = *media

		switch search["Media"] {
			case "twitter":
				search["Query"] = *query
				search["Lang"] = "pt"
				search["Geocode"] = ""
				// search["__resource"] = `{"klout_api_key":"cxyaz95kx93ep6bf4zzt6676"}` // fabriciojs
			case "1_2":
			case "1_1": {
				// Twitter DMs retrieving
			}
			case "youtube":
				search["Query"] = *query
				search["Lang"] = "pt"
			case "3":
				search["Query"] = *query
				// search["Locale"] = "pt_BR"
				search["Locale"] = ""
			case "flickr":
				search["Query"] = *query
			case "facebook":
				search["Query"] = *query
				search["Locale"] = "pt_BR"
				search["Until"] = ""
			case "6":
				// Orkut
				search["Settings"] = `{"Cmm":["4249"]}` // Brasil
				// search["Settings"] = `{"Cmm":["33383"]}` // Unesp
				search["Query"] = *query
			case "7":
				search["Query"] = *query
				search["Locale"] = "br"
			case "google_plus":
				search["Query"] = *query
				search["Lang"] = "pt"
			case "rss":
				search["Url"] = "{\"Urls\":[\"http://rss.terra.com.br/0,,EI306,00.xml\",\"http://rss.terra.com.br/0,,EI14416,00.xml\"]}"
				search["Query"] = *query
			case "facebook_pages":
				// FB pages
				search["Id"] = *query
			case "foursquare":
				// Foursquare
				search["TypeSearch"] = `{"types":["tips","photos"]}`
				search["Venue"] = "4d29b8ef3c795481133ada9b"
			case "linkedin":
				// Linkedin
				search["Settings"] = `{"GroupId":["3937610"]}`
			case "slideshare":
				// Slideshare
				search["Query"] = "Plano"
			case "reclame_aqui":
				// search["Query"] = "Plano"
				search["Id"] = "7712"
			case "instagram":
				search["Query"] = *query
		}

		jsonBytes, _ := json.Marshal(search)
		*searchjson = string(jsonBytes)
	} else {
		err := json.Unmarshal([]byte(*searchjson), &search)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	for {
		// conn, err := net.Dial("tcp", "localhost:" + )
		// fmt.Println("75.126.48.226:" + *search_port)
		conn, err := net.Dial("tcp", "75.126.48.226:" + *search_port)
		// conn, err := net.Dial("tcp", "localhost:" + *search_port)

		if err != nil {
			fmt.Printf("Could not connect: %s\n", err)
			time.Sleep(time.Second)
			continue
		}

		client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

		var responseString string

		client.Call("RPCSender.AddExternalSearch", searchjson, &responseString)

		responseData := strings.Split(responseString, ";")

		searchSpeed, _ := strconv.Atoi(responseData[2])

		fmt.Println(responseString)
		fmt.Println(responseData)

		search["LastItem"] = responseData[1]
		search["First"] = "0"

		fmt.Printf("\n%s - sleep: %v\n", search["Extid"], time.Second * time.Duration(searchSpeed))

		time.Sleep(time.Second * time.Duration(searchSpeed))
	}
}

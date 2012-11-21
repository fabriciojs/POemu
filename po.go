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

var media = flag.String("media", "", "Media Code [1, 14]")
var query = flag.String("query", "", "Keyword to query")
var limit = flag.String("limit", "2000", "limit items amount, first import")
var search_port = flag.String("rpc_port", "8222", "The RPC port to connect to the search service")

var cod_busca = flag.String("cod_busca", "", "o Código da busca")
var cod_monitoramento = flag.String("cod_monitoramento", "", "o Código do monitoramento")
var cod_conta = flag.String("cod_conta", "", "o Código da conta")

func main() {
	flag.Parse()

	if *query == "" {
		fmt.Println("You must provide the --query param.")
		os.Exit(1)
	}

	if *cod_conta == "" {
		fmt.Println("You must provide the --cod_conta param.")
		os.Exit(1)
	}

	if *cod_monitoramento == "" {
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

	search := make(map[string]string)

	search["LastItem"] = ""
	search["FormattedResult"] = "0"
	search["First"] = "1"
	search["Feed"] = "http://livebuzz.com.br/dev/importapi/search/store"
	// search["Extid"] = fmt.Sprintf("116.116.0.%s.%s", *cod_busca, *media)
	search["Extid"] = fmt.Sprintf("%s.%s.0.%s.%s", *cod_conta, *cod_monitoramento, *cod_busca, *media)

	search["Limit"] = *limit

	search["__ip"] = "192.168.0.138"

	search["Media"] = *media

	for {
		conn, err := net.Dial("tcp", "localhost:" + *search_port)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not connect: %s\n", err)
			os.Exit(1)
		}

		client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

		switch search["Media"] {
			case "twitter":
				search["Query"] = *query
				search["Lang"] = "pt"
				search["Geocode"] = ""
				// search["__resource"] = `{"klout_api_key":"cxyaz95kx93ep6bf4zzt6676"}` // fabriciojs
				search["__resource"] = `{"klout_api_key":"9ftwwawy88z434rxw7es8w2u"}` // livebuzz
			case "1_2":
			case "1_1": {
				// Twitter DMs retrieving
				search["__resource"] = `{"consumer_key":"d7a1m12fUC8kEKpw9vfXCA","consumer_secret":"WK8TE5tvCoasHl7CosvBqCgiMwDAzVkScevWnB6ZQ","token":"465451251-hhYKSIWJ89475CnzW7iTPIv0QwtSkW42U3WYJZ3c","token_secret":"yxo3lunJFpqWXsQ4DbT8fSEoy1tA7aGNfNr3U7RDsI"}`
			}
			case "youtube":
				search["Query"] = *query
				search["Lang"] = "pt"
				search["__key1"] = "AI39si6gHpLx9SQEZfUdKv_LvQl-esjOzC-g2FPgNQkbgmQYjHTJGP7GVc6lXqNZ87qPkCKLMijJ3O87yDMbFcr5t0aXEwr9hw"
			case "3":
				search["Query"] = *query
				// search["Locale"] = "pt_BR"
				search["Locale"] = ""
			case "flickr":
				search["Query"] = *query
				search["__resource"] = `{"key":"7e2d07aeeb4c5d4cd41b6776d8b0029e","secret":"d21bdf9d05db4e06"}`
			case "facebook":
				search["__resource"] = "150062421699542|1ac6f54d89f75603d807c82cb59bd2f7"
				search["Query"] = *query
				search["Locale"] = "pt_BR"
				search["Until"] = ""
				// search["Until"] = "1344729600"
			case "6":
				// Orkut
				search["Settings"] = `{"Cmm":["4249"]}` // Brasil
				// search["Settings"] = `{"Cmm":["33383"]}` // Unesp
				search["Query"] = *query
				search["__key1"] = `livebuzzdev@gmail.com;goliv123`
			case "7":
				search["Query"] = *query
				search["Locale"] = "br"
			case "google_plus":
				search["Query"] = *query
				search["__key1"] = "AIzaSyCyPp5S-_gCnwNDjrBeWwSLf9-2nXj0e58"
				search["Lang"] = "pt"
			case "rss":
				search["Url"] = "{\"Urls\":[\"http://rss.terra.com.br/0,,EI306,00.xml\",\"http://rss.terra.com.br/0,,EI14416,00.xml\"]}"
				search["Query"] = *query
			case "facebook_pages":
				// FB pages
				search["__key1"] = "150062421699542|1ac6f54d89f75603d807c82cb59bd2f7"
				search["id"] = *query
			case "foursquare":
				// Foursquare
				// search["TypeSearch"] = "{\"types\":[\"tips\",\"photos\"]}"
				search["TypeSearch"] = `{"types":["tips","photos"]}`
				search["__resource"] = `{"client_id":"V2MICHIQ3KTMDGMPN350TRNBTBHMCFACCBGH1PII21SMLI3X","client_secret":"W42XRVPWYGTOZTF1AYMQ3E4IQREWL5L032A0O0XNHSBS0KWX"}`
				search["Venue"] = "4d29b8ef3c795481133ada9b"
			case "linkedin":
				// Linkedin
				search["__resource"] = `{"consumer_key":"w73iazuam3iu","consumer_secret":"6nWu3USoO3OVQRQt","token":"1de5ebf3-53e7-480f-8b4d-0e25fecddc19","token_secret":"6b36f0f7-91c0-486f-98b1-6e6e98c25fa7"}`
				search["Settings"] = `{"GroupId":["3937610"]}`
			case "slideshare":
				// Slideshare
				search["__resource"] = `{"key":"8dguR2ZP","secret":"1PbNlblE"}`
				search["Query"] = "Plano"
			case "reclame_aqui":
				// search["Query"] = "Plano"
				search["id"] = "7712"
			case "instagram":
				search["__resource"] = `{"client_id":"3ffba910b8d044a38923f95f81757bf9"}`
				search["Tag"] = "itau"
		}

		jsonBytes, _ := json.Marshal(search)
		jsonString := string(jsonBytes)

		var responseString string

		client.Call("RPCSender.AddExternalSearch", &jsonString, &responseString)

		// fmt.Printf("%s : %s\n", search["Extid"], responseString)

		responseData := strings.Split(responseString, ";")

		searchSpeed, _ := strconv.Atoi(responseData[4])

		search["LastItem"] = responseData[4]
		search["First"] = "0"

		fmt.Printf("%s - sleep: %d minute(s)\n", search["Extid"], searchSpeed / 60)

		time.Sleep(time.Second * time.Duration(searchSpeed))
	}
}

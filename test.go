package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	date := "2012-10-31T17:19:07.005Z"

	date = strings.Replace(date, "Z", "", -1)

	fmt.Println(date)

	t, e := time.Parse("2006-01-02T15:04:05", date)

	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(t.String())
	}
}

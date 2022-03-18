package main

import (
	"fmt"
	"mtsbank_golang/rsk_parsing/site"
	"time"
)

func main() {

	params := map[string]string{
		"status":      "all",
		"gov":         "28",
		"paginate_by": "50",
		"page":        "4",
	}

	parser := site.NewRskParser()

	siteWP1 := site.NewRskSite(params)
	cp1 := site.NewConcurrentParser(parser, siteWP1)

	in := make(chan site.Result, 3)

	go cp1.ParseAll(in)

	start := time.Now()

	for r := range in {
		if r.Err != nil {
			fmt.Println("ERROR ", r.Err)
		} else {
			fmt.Println(r.Bws.String())
		}
	}

	fmt.Println(time.Now().Sub(start))
}

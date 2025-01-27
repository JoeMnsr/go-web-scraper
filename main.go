package main

import (
	"log"
)

func main() {
	url := "https://joemansour.vercel.app"

	// fetch data
	resp, err := fetchPage(url)
	if err != nil {
		log.Fatalf("Error fetching page %s", err)
	}

	// parse the html to extract data
	parseHTML(resp)
}

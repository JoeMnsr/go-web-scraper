package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func parseHTML(resp *http.Response) {
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Failed to fetch page: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Error parsing HTML %s", err)
	}

	// iterate through each paragraph and return it's content
	doc.Find("p").Each(func(index int, element *goquery.Selection) {
		headline := element.Text()
		fmt.Println(headline)
	})
}

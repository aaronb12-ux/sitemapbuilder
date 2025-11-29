package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"aaron.com/sitemapbuilder/bfs"
	"github.com/aaronb12-ux/linkparserhtml"
)

func getHTML(url string) []byte {
	

	response, err := http.Get(url)

	if err != nil {
		log.Fatal("error fetching URL: ", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal("error reading response body: ", err)
	}

	return body 
}

func BFS() []string {

	
}

func main() {

	var url string

	flag.StringVar(&url, "url", "https://www.calhoun.io/", "the url to extract links from")
	
	flag.Parse()

	ans := BFS()

	var body []byte = getHTML(url)

	links := linkparserhtml.GetLinks(strings.NewReader(string(body)))

	


	
 
}
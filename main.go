package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"aaron.com/sitemapbuilder/bfs"
	//"strings"
	//"aaron.com/sitemapbuilder/bfs"
	//"github.com/aaronb12-ux/linkparserhtml"
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


func main() {

	var url string

	flag.StringVar(&url, "url", "https://example.com", "the url to extract links from")
	
	flag.Parse()

	ans := bfs.Bfs(url)

	fmt.Println(ans)


	
 
}
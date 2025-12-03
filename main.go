package main

import (
	"flag"
	"fmt"
	"aaron.com/sitemapbuilder/bfs"
)


func main() {

	var url string

	flag.StringVar(&url, "url", "https://includedavis.com/", "the url to extract links from")
	
	flag.Parse()

	ans := bfs.Bfs(url)

	fmt.Println("All the valid links are", ans)

}

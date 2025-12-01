package bfs

import (
	"fmt"
	"net/http"
	"strings"
	"io"
	"github.com/aaronb12-ux/linkparserhtml"
	"log"
	"net/url"
)

type Queue struct {
	urls []string
}

func (q *Queue) Enqueue(url string) {
	q.urls = append(q.urls, url)
}

func (q *Queue) Dequeue() string {
	if len(q.urls) == 0 {
		fmt.Println("queue is empty")
		return "empty"
	}
	front := q.urls[0]
	q.urls = q.urls[1:]
	return front
}

func CreateQueue() *Queue {
	queue := &Queue{}
	return queue
}

func makeAbsolute(currentPath string, baseUrl string ) string {

	base, err := url.Parse(baseUrl)

	if err != nil {
		return ""
	}

	ref, err := url.Parse(currentPath)
	if err != nil {
		return ""
	}

	return base.ResolveReference(ref).String()

}

func GetHTML(url string, baseurl string) []byte {

	validUrl := makeAbsolute(url, baseurl)

	response, err := http.Get(validUrl)

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


func validLink(link string) bool {
	//checks if a link is in the domain
	if strings.HasPrefix(link, "./") {
		return true
	}
	return false
}

func Bfs(baseurl string) []string { //run a bfs from the root (base domain url)

	var ans []string 

	//create queue and initialize it
	queue := CreateQueue() 
	queue.Enqueue(baseurl)

	//create seenSet and initilize it
	seenUrls := make(map[string]bool)
	seenUrls[baseurl] = true

	for len(queue.urls) > 0 {

		currentURL := queue.Dequeue()

		var body []byte = GetHTML(currentURL, baseurl)

		links := linkparserhtml.GetLinks(strings.NewReader(string(body)))

		//now get the valid links only...
		for _ , link := range links {
			
			if !seenUrls[link.Href] && validLink(link.Href) {
				ans = append(ans, link.Href)
				seenUrls[link.Href] = true
				queue.Enqueue(link.Href)
			} 
		}
	}
	
	return ans
}

package bfs

import (
	"fmt"
	"net/http"
	"strings"
	"io"
	"github.com/aaronb12-ux/linkparserhtml"
	"log"
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
	return front
}

func CreateQueue() *Queue {
	queue := &Queue{}
	return queue
}


func GetHTML(url string) []byte {


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

func validLink(link string) bool {
	//checks if a link is in the domain
	if link[0] == '/' {
		return true
	}
	return false
}

func Bfs(url string) []string { //run a bfs from the root (base domain url)

	var ans []string 

	//create queue and initialize it
	queue := CreateQueue() 
	queue.Enqueue(url)


	//create seenSet and initilize it
	seenUrls := make(map[string]bool)
	seenUrls[url] = true


	for len(queue.urls) > 0 {

		currentURL := queue.Dequeue()

		var body []byte = GetHTML(currentURL)

		links := linkparserhtml.GetLinks(strings.NewReader(string(body)))

		fmt.Println(links)

		//now get the valid links only...
		for _ , link := range links {
			if !seenUrls[link.Href] && validLink(link.Href) {
				ans = append(ans, link.Href)
				seenUrls[link.Href] = true
				queue.Enqueue(link.Href)
			} else {
				continue
			}
		}
	}
	
	return ans
}



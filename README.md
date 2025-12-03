# Web Crawler - BFS Site Mapper

A simple web crawler built in Go that uses Breadth-First Search (BFS) to discover and map all internal links on a website.

## What It Does

Starting from a base URL, the crawler:
- Fetches the HTML of each page
- Extracts all links from the page
- Follows internal links (relative paths like `./about` or `/contact`)
- Avoids revisiting pages using a seen set
- Returns a list of all discovered URLs

## Key Learning Points

- **BFS Implementation**: Applied BFS algorithm to a real-world problem (web crawling)
- **HTML Parsing**: Used Go's `golang.org/x/net/html` package to parse HTML
- **URL Handling**: Learned to resolve relative URLs to absolute URLs
- **Queue Data Structure**: Implemented a custom queue for BFS traversal
- **HTTP Requests**: Made GET requests and processed responses

## How It Works

1. Start with a base URL and add it to the queue
2. Dequeue a URL and fetch its HTML
3. Parse the HTML and extract all links
4. Filter for valid internal links
5. Add new unseen links to the queue
6. Repeat until queue is empty

## Usage
```go
package main

import (
    "fmt"
    "yourmodule/bfs"
)

func main() {
    baseURL := "https://example.com"
    links := bfs.Bfs(baseURL)
    
    fmt.Println("Found links:")
    for _, link := range links {
        fmt.Println(link)
    }
}
```

## Project Structure

- `bfs/`: Main crawler logic with BFS implementation
- `linkparserhtml/`: HTML parsing and link extraction

## Notes

This was a learning project focused on applying BFS to a practical problem rather than just LeetCode-style graph problems. The goal was to understand the algorithm in a real-world context.

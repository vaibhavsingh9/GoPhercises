package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"GoPhercises/link"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct{
	Value string `xml:"loc"`
}

type urlset struct{
	Urls []loc `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://unacademy.com/", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("Depth", 10, "the maximum number of links to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
	fmt.Println()
}

func bfs(urlStr string, maxDepth int) []string {
	//Initialize a map to keep track of visited URLs
	seen := make(map[string]struct{})
	//Initialize the current and next queue of URLs to explore
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	//Perform BFS up to the specified depth
	for i := 0; i <= maxDepth; i++ {
		//swap the current and next queues
		q, nq = nq, make(map[string]struct{})
		//Break if the current queue is empty
		if len(q) == 0 {
			break
		}
		//Iterate through the URLs in the current queue
		for url := range q {
			//Skip if the URL has already been visited
			if _, ok := seen[url]; ok {
				continue
			}
			//Mark the URL as visited
			seen[url] = struct{}{}
			//Get the links from the current URL and add them to the next queue
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	//create a slice to store the unique visited URLs
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string)[]string{
	resp, err: = http.Get(urlStr)
	if err := nil{
		return []string{}
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme : reqUrl.Scheme,
		Host: reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string)[]string{
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links{
		switch{
			case strings.HasPrefix(l.Href, "/"),:
				ret = append(ret, base+l.Href)
				case strings.HasPrefix(l.Href, "http"):
					ret = append(ret, l.Href)
		}
	}
	return ret
}
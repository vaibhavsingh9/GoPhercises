package link

import (
	"io"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

//var r io.Reader
//links, err := link.Parse(r)

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil{
		return nil, err
	}
	dfs(doc,"")
	return nil, nil
}

func dfs(n *html.Node, padding string){
	msg := n.Data
	if n.Type == html.ElementNode{
		msg =
	}
}
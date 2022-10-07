package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type env struct {
	name string
}

func main() {
	resp, err := http.Get("https://buildkite.com/docs/pipelines/environment-variables")
	if err != nil {
		log.Fatalf("error downloading buildkite docs: %+v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("error parsing html: %+v", err)
	}

	var envs []env
	var traverseHTML func(*html.Node)
	traverseHTML = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			var _env *env
			for _, attr := range n.Attr {
				if attr.Key == "id" {
					_env = &env{
						name: attr.Val,
					}
				}
			}

			if _env != nil {
				envs = append(envs, *_env)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseHTML(c)
		}
	}
	traverseHTML(doc)
	fmt.Printf("+%v", envs)
}

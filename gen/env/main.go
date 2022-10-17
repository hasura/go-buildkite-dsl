//go:generate go run main.go

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type env struct {
	Name string
}

// TODO handle BUILDKITE_AGENT_META_DATA_

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
						Name: attr.Val,
					}
				}
			}

			if _env != nil {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "td" {
						// fmt.Printf("%+v\n", c.Attr)
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

	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalf("unable to get wd: %+v", err)
	// }
	// fmt.Printf("wd: %s", wd)

	file, err := os.OpenFile("../../env/env.go", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("unable to open file: %+v", err)
	}

	envTemplate, err := template.New("env.go.tmpl").ParseFiles("env.go.tmpl")
	if err != nil {
		log.Fatalf("error creating template: %+v", err)
	}
	err = envTemplate.Execute(file, envs)
	if err != nil {
		log.Fatalf("error executing template: %+v", err)
	}
}

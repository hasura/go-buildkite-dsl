//go:generate go run main.go

package main

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type env struct {
	Name        string
	Description string
}

// TODO handle BUILDKITE_AGENT_META_DATA_
var envsToExclude = []string{
	"BUILDKITE_AGENT_META_DATA_",
}

func main() {
	resp, err := http.Get("https://buildkite.com/docs/pipelines/environment-variables")
	if err != nil {
		log.Fatalf("error downloading buildkite docs: %+v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("error parsing html: %+v", err)
	}

	var envs []env
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		var envVar *env
		if envName, exists := s.Attr("id"); exists {
			for _, envToExclude := range envsToExclude {
				if envName == envToExclude {
					return
				}
			}

			envVar = &env{
				Name: envName,
			}
		}

		if envVar == nil {
			return
		}
		envVar.Description = s.Find("td").Text()
		envVar.Description = commentify(envVar.Description)

		envs = append(envs, *envVar)
	})

	os.RemoveAll("../../env/env.go")

	file, err := os.OpenFile("../../env/env.go", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("unable to open file: %+v", err)
	}

	envTemplate, err := template.New("env.go.tmpl").Funcs(
		template.FuncMap{
			"unescape": func(s string) string {
				return html.UnescapeString(s)
			},
		},
	).ParseFiles("env.go.tmpl")
	if err != nil {
		log.Fatalf("error creating template: %+v", err)
	}

	var data []byte
	buffer := bytes.NewBuffer(data)
	err = envTemplate.Execute(buffer, envs)
	if err != nil {
		log.Fatalf("error executing template: %+v", err)
	}

	// text/template package is auto-escaping some HTML characters like double-quotes
	// hence unescaping before we write it to the file
	newContent := html.UnescapeString(buffer.String())
	file.WriteString(newContent)
	fmt.Printf("generated %s\n", file.Name())
}

func commentify(comment string) string {
	comment = strings.TrimSpace(comment)
	comment = strings.ReplaceAll(comment, "\n\n", "")
	var commentedLines []string
	for _, line := range strings.Split(comment, "\n") {
		commentedLines = append(commentedLines, fmt.Sprintf("// %s", strings.Trim(line, " ")))
	}
	return strings.Join(commentedLines, "\n")
}

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://buildkite.com/docs/pipelines/environment-variables")
	if err != nil {
		log.Fatalf("error downloading buildkite docs: %+v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error while reading response: %+v", err)
	}

	log.Printf("body: %s", string(body))
}

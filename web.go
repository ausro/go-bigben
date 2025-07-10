package main

import (
	"io"
	"log"
	"net/http"
)

func get(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to GET: %s", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatalf("Failed to read body: %s", err)
	}

	return string(body)
}

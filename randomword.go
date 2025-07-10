package main

import (
	"log"
	"regexp"
)

const (
	WordURL = "https://randomword.com/"
)

type Word struct {
	Name       string
	Definition string
}

func getRandomWord() Word {
	body := get(WordURL)

	word := getWordFromBody(body)
	log.Printf("Got word: %s with definition: %s", word.Name, word.Definition)

	return word
}

func getWordFromBody(body string) Word {
	// <div id="random_word">[word]</div> ... <div id="random_word_definition">[definition]</div>
	r := regexp.MustCompile("id=\"random_word\">(?P<Name>[A-Za-z ,;-]+)<.+[\n].+id=\"random_word_definition\">(?P<Definition>[A-Za-z ,;-]+)")

	names := r.FindStringSubmatch(body)
	if names == nil {
		log.Println("No word found in response body.")
	}

	return Word{
		Name:       names[1],
		Definition: names[2],
	}
}

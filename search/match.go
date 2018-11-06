package search

import (
	"log"
)

type Result struct {
	Field string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, err)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

func Display(reulsts cha *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
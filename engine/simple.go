package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

// input many seed urls
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	// enqueue the seed urls
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// while queue is not empty
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)

		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)

		// handle the data
		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

// merge Fetcher and Parser into this function
func worker(r Request) (ParserResult, error) {
	// get the page content of queue.pop()
	log.Printf("Fetching url: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error occured at url %s: %v", r.Url, err)
		return ParserResult{}, err
	}

	// use the specific ParserFunc to parse new urls and enqueue
	return r.ParserFunc(body), nil
}
package engine

import (
	"log"

	"baseLearn/15.爬虫实现--单任务版-并发版/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+
			"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}

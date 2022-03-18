package site

import (
	"sync"
)

type Result struct {
	Err error
	Bws BlockedWebsite
}

type ConcurrentParser struct {
	parser BlockedWebsitesParser
	siteWP SiteWithPages
}

func NewConcurrentParser(parser BlockedWebsitesParser, siteWP SiteWithPages) *ConcurrentParser {
	return &ConcurrentParser{parser: parser, siteWP: siteWP}
}

func (s *ConcurrentParser) ParseAll(in chan<- Result) {

	iter, err := s.siteWP.CreatePageIterator()

	if err != nil {
		return
	}

	wg := &sync.WaitGroup{}

	for iter.Next() {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			url := s.siteWP.ConstructUrl(page)
			blockedWebsites, err1 := s.parser.Parse(url)

			if err1 != nil {
				in <- Result{Err: err1}
				return
			}

			for _, bws := range blockedWebsites {
				in <- Result{Bws: bws}
			}

		}(iter.Current())
	}

	wg.Wait()
	close(in)

	return
}

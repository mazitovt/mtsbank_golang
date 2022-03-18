package site

import (
	"errors"
	"net/url"
	"strconv"
	"sync"
)

const (
	domain = "https://reestr.rublacklist.net/?"
)

type SiteWithPages interface {
	CreatePageIterator() (Iterator, error)
	SetParam(string, string)
	ConstructUrl(int) string
}

type RskSite struct {
	params url.Values
	rwMtx  sync.RWMutex
}

func NewRskSite(params map[string]string) *RskSite {

	urlValues := url.Values{}

	for k, v := range params {
		urlValues[k] = []string{v}
	}

	return &RskSite{params: urlValues}
}

func (s *RskSite) SetParam(key string, value string) {
	s.rwMtx.Lock()
	defer s.rwMtx.Unlock()
	s.params.Set(key, value)
}

func (s *RskSite) CreatePageIterator() (iter Iterator, err error) {

	s.rwMtx.RLock()
	defer s.rwMtx.RUnlock()

	if !s.params.Has("page") {
		err = errors.New("no param 'page'")
		return
	}

	finishPage, err := strconv.ParseInt(s.params.Get("page"), 10, 64)

	if err != nil {
		return
	}

	iter = NewPageIterator(1, int(finishPage), s.ConstructUrl)

	return
}

func (s *RskSite) ConstructUrl(page int) string {
	s.rwMtx.Lock()
	defer s.rwMtx.Unlock()

	old := s.params.Get("page")
	defer s.params.Set("page", old)

	s.params.Set("page", strconv.FormatInt(int64(page), 10))

	res := domain + s.params.Encode()

	return res
}

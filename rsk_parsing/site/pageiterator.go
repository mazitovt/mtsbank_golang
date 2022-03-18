package site

import "errors"

type Iterator interface {
	Next() bool
	Current() int
	Err() error
}

type PageIterator struct {
	cur          int
	max          int
	constructUrl func(int) string
	err          error
}

var (
	startPageErr  = errors.New("startPage less than -1")
	finishPageErr = errors.New("finishPage less than startPage")
)

func NewPageIterator(startPage int, finishPage int, constructUrl func(int) string) (iter *PageIterator) {

	var err error

	if startPage < -1 {
		err = startPageErr
	}

	if finishPage < startPage {
		err = finishPageErr
	}

	iter = &PageIterator{
		cur:          startPage - 1,
		max:          finishPage,
		constructUrl: constructUrl,
		err:          err,
	}

	return
}

func (iter *PageIterator) Next() bool {
	if iter.err != nil {
		return false
	}
	iter.cur++
	return iter.cur <= iter.max
}

func (iter *PageIterator) Current() int {
	return iter.cur
}

func (iter *PageIterator) Err() error {
	return iter.err
}

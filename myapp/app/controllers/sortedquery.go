package controllers

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
)

type SortedQuery struct {
	_map map[string]string
	keys []string
}

func (sq *SortedQuery) Len() int {
	return len(sq._map)
}

func (sq *SortedQuery) Less(i, j int) bool {
	return sq.keys[i] < sq.keys[j]
}

func (sq *SortedQuery) Swap(i, j int) {
	sq.keys[i], sq.keys[j] = sq.keys[j], sq.keys[i]
}

func (sq *SortedQuery) String() string {

	sort.Sort(sq)
	values := make([]string, len(sq.keys))
	for i, key := range sq.keys {
		values[i] = fmt.Sprintf("%s=%s", key, url.QueryEscape(sq._map[key]))
	}
	return strings.Join(values, "&")
}

func NewSortedQuery(m map[string]string) *SortedQuery {
	sq := new(SortedQuery)
	sq._map = m
	sq.keys = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sq.keys[i] = key
		i++
	}
	return sq
}
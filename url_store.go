package main

import (
	"strconv"
	"sync/atomic"
)

type UrlStore struct {
	id   uint64
	urls map[string]string
}

func NewUrlStore() *UrlStore {
	return &UrlStore{
		id:   0,
		urls: make(map[string]string),
	}
}

func (store *UrlStore) expand(token string) string {
	return store.urls[token]
}

func (store *UrlStore) shorten(url string) string {
	token := store.nextToken()
	store.urls[token] = url
	return token
}

func (store *UrlStore) nextId() uint64 {
	return atomic.AddUint64(&store.id, 1)
}

func (store *UrlStore) nextToken() string {
	id := store.nextId()
	return store.tokenize(id)
}

func (store *UrlStore) tokenize(id uint64) string {
	return strconv.FormatUint(id, 10)
}

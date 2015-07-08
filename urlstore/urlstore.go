package urlstore

import (
	"strconv"
	"sync/atomic"
)

// URLStore represents the complete state of the world.
type URLStore struct {
	id   uint64
	urls map[string]string
}

// NewURLStore returns a new/empty URLStore
func NewURLStore() *URLStore {
	return &URLStore{
		id:   0,
		urls: make(map[string]string),
	}
}

// Expand fetches a URL from the store, or returns nothing if not found
func (store *URLStore) Expand(token string) string {
	return store.urls[token]
}

// Shorten stores a URL and returns a token
func (store *URLStore) Shorten(url string) string {
	token := store.nextToken()
	store.urls[token] = url
	return token
}

func (store *URLStore) nextID() uint64 {
	return atomic.AddUint64(&store.id, 1)
}

func (store *URLStore) nextToken() string {
	id := store.nextID()
	return store.tokenize(id)
}

func (store *URLStore) tokenize(id uint64) string {
	return strconv.FormatUint(id, 10)
}

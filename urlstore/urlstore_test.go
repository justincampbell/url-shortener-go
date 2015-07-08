package urlstore

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestShorten(t *testing.T) {
	urlStore := NewURLStore()
	url := "http://justincampbell.me"
	token := "1"
	assert.Equal(t, token, urlStore.Shorten(url))
}

func TestExpand(t *testing.T) {
	urlStore := NewURLStore()
	url := "http://justincampbell.me"
	token := "1"
	urlStore.urls[token] = url
	assert.Equal(t, url, urlStore.Expand(token))
}

func TestNextId(t *testing.T) {
	urlStore := NewURLStore()
	assert.Equal(t, uint64(1), urlStore.nextID())
	assert.Equal(t, uint64(2), urlStore.nextID())
}

func TestNextToken(t *testing.T) {
	urlStore := NewURLStore()
	assert.Equal(t, "1", urlStore.nextToken())
	assert.Equal(t, "2", urlStore.nextToken())
}

func TestTokenize(t *testing.T) {
	urlStore := NewURLStore()
	assert.Equal(t, "1", urlStore.tokenize(1))
}

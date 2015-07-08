package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestShorten(t *testing.T) {
	urlStore := NewUrlStore()
	url := "http://justincampbell.me"
	token := "1"
	assert.Equal(t, token, urlStore.shorten(url))
}

func TestExpand(t *testing.T) {
	urlStore := NewUrlStore()
	url := "http://justincampbell.me"
	token := "1"
	urlStore.urls[token] = url
	assert.Equal(t, url, urlStore.expand(token))
}

func TestNextId(t *testing.T) {
	urlStore := NewUrlStore()
	assert.Equal(t, uint64(1), urlStore.nextId())
	assert.Equal(t, uint64(2), urlStore.nextId())
}

func TestNextToken(t *testing.T) {
	urlStore := NewUrlStore()
	assert.Equal(t, "1", urlStore.nextToken())
	assert.Equal(t, "2", urlStore.nextToken())
}

func TestTokenize(t *testing.T) {
	urlStore := NewUrlStore()
	assert.Equal(t, "1", urlStore.tokenize(1))
}

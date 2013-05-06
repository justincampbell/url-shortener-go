package main

import (
  "github.com/bmizerany/assert"
  "testing"
)

func reset() {
  id = 0
  urls = make(map[string]string)
}

func TestShorten(t *testing.T) {
  reset()
  url := "http://justincampbell.me"
  token := "1"
  assert.Equal(t, token, shorten(url))
}

func TestExpand(t *testing.T) {
  url := "http://justincampbell.me"
  token := "1"
  urls[token] = url
  assert.Equal(t, url, expand(token))
}

func TestNextId(t *testing.T) {
  reset()
  assert.Equal(t, uint64(1), nextId())
  assert.Equal(t, uint64(2), nextId())
}

func TestNextToken(t *testing.T) {
  reset()
  assert.Equal(t, "1", nextToken())
  assert.Equal(t, "2", nextToken())
}

func TestTokenize(t *testing.T) {
  assert.Equal(t, "1", tokenize(1))
}

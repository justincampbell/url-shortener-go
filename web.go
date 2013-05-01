package main

import (
  "flag"
  "fmt"
  "net/http"
  "strconv"
)

var (
  id   = 0
  port = flag.String("port", "8080", "port to listen on")
  urls = make(map[string]string)
)

func init() {
  flag.Parse()
}

func main() {
  http.HandleFunc("/", expandHandler)
  http.HandleFunc("/shorten", shortenHandler)
  fmt.Println("Listening on", *port)
  http.ListenAndServe(":"+*port, nil)
}

func expandHandler(response http.ResponseWriter, request *http.Request) {
  fmt.Println(request.Method, request.RequestURI, request.RemoteAddr)

  if request.RequestURI == "/" {
    url := "http://github.com/justincampbell/url-shorteners"
    http.Redirect(response, request, url, 301)
    return
  }

  token := request.URL.Path[len("/"):]
  url := expand(token)

  if url == "" {
    http.NotFound(response, request)
    return
  }

  http.Redirect(response, request, url, 301)
}

func shortenHandler(response http.ResponseWriter, request *http.Request) {
  fmt.Println(request.Method, request.RequestURI, request.RemoteAddr)
  uri := request.RequestURI
  parameter := "/shorten?url="

  if len(uri) < len(parameter) {
    http.Error(response, "Bad Request", 400)
    return
  }

  url := uri[len(parameter):]

  if url == "" {
    http.Error(response, "Bad Request", 400)
    return
  }

  token := shorten(url)
  fmt.Fprintf(response, "/%s", token)
}

func expand(token string) string {
  return urls[token]
}

func shorten(url string) string {
  token := nextToken()
  urls[token] = url
  return token
}

func nextId() int {
  id = id + 1
  return id
}

func nextToken() string {
  id := nextId()
  return tokenize(id)
}

func tokenize(id int) string {
  return strconv.Itoa(id)
}

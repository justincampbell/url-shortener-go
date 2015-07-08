package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/justincampbell/url-shortener-go/url_store"
)

var (
	idRequest  = make(chan bool)
	idResponse = make(chan int)
	port       = flag.String("port", "8080", "port to listen on")
	urlStore   = url_store.NewUrlStore()
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
	url := urlStore.Expand(token)

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

	token := urlStore.Shorten(url)
	fmt.Fprintf(response, "/%s", token)
}

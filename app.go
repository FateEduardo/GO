package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var INDEX_HTML []byte

func main() {
	http.HandleFunc("/", IndexHandler)
	log.Println("Listening...")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(INDEX_HTML)

}

func init() {
	INDEX_HTML, _ = ioutil.ReadFile("./tmpl/index.html")

}

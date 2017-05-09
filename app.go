package hello

import (
	"io/ioutil"
	"log"
	"net/http"
)

var INDEX_HTML []byte



func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(INDEX_HTML)

}

func init() {
	http.HandleFunc("/", IndexHandler)
	log.Println("Listening...")
	INDEX_HTML, _ = ioutil.ReadFile("./tmpl/index.html")

}

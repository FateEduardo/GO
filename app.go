package main

import (
  "log"
  "net/http"
  "io/ioutil"
)

var INDEX_HTML []byte


func main() {
  http.HandleFunc("/", IndexHandler)
  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
    w.Write(INDEX_HTML)
   
}

func init(){
    INDEX_HTML, _ =ioutil.ReadFile("./tmpl/index.html")
   
}
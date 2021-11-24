package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h2>Air自动重载<h2>")
}

func main(){
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3020", nil)
}
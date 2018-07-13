package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("assigning handlers")
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(resw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resw, "this is just a message passed at %s", time.Now().String())
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
   
    if err := http.ListenAndServe(":11130", nil); err != nil {
        log.Fatal(err)
	}
	
}

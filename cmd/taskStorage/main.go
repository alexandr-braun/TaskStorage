package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello, %s !", r.URL.Path)
}

func main() {
	NewDatabaseRegistrar().addDatabase()

	http.HandleFunc("/", handler)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	select {}
}

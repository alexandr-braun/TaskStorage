package main

import (
	"fmt"
	"log"
	"net/http"
)

// TODO init query handlers separately
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello, %s !", r.URL.Path)
}

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	NewDatabaseRegistrar().connectToDatabase(cfg)

	// TODO init presentation layer
	http.HandleFunc("/", handler)
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	select {}
}

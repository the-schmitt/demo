package main

import (
	"fmt"
	"log"
	"net/http"

	levi "github.com/schmitt/demo-mono/import-levi"
)

func main() {
	// Use the import-levi package
	levi.PrintHello()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		levi.PrintHello()
		fmt.Fprintf(w, "Hello from fake-server!")
	})

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

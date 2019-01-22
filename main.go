package main

import (
	"fmt"
	"net/http"
	"os"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/test", testHandler)

	fmt.Printf("[info] test\n")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		os.Exit(1)
	}
}

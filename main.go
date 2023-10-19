package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "We did it..... 🎉 ")
	})
	fmt.Println("Server has started 🏃")
	http.ListenAndServe(":8080", nil)
}

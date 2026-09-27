package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	host, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from %s\n", host)
	})
	http.ListenAndServe(":8080", nil)
}

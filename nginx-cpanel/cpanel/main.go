package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/cpanel", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login to CPanel page Served by Go"))
	})

	fmt.Println("Starting CPanel server at :2083")
	err := http.ListenAndServe(":2083", nil)
	if err != nil {
		log.Fatalf("failed to start %v", err)
	}
}

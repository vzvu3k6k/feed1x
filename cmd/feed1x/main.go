package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vzvu3k6k/feed1x"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT is not found")
	}

	log.Fatal(http.ListenAndServe(":"+port, feed1x.NewServer()))
}

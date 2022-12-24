package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1]
	fs := http.FileServer(http.Dir(os.Args[2]))
	log.Fatal(http.ListenAndServe(":"+port, fs))
}

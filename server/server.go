package main

import (
	"fmt"
	"log"
	"net/http"
)

func getLevel(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "the level will be provided here \n")
	fmt.Println("level was requested and provided")
}

func main() {
	fmt.Println("Running the server...")

	http.HandleFunc("/level", getLevel)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

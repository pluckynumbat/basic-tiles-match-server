package main

import (
	"fmt"
	"net/http"
)

func getLevel(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "the level will be provided here \n")
	fmt.Println("level was requested and provided")
}

func main() {
	fmt.Println("Server Coming Soon...")

	http.HandleFunc("/level", getLevel)

	http.ListenAndServe(":8090", nil)
}

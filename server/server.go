package main

import (
	"fmt"
	"log"
	"net/http"
)

const levelString = `{
  "name": "From The Server",
  
  "seed": 10,
  
  "colorCount": 4,

  "colorPalette": ["R","G","B","V"],
  
  "gridLength": 5,
  
  "isStartingGridFixed": true,
  
  "startingGrid": 
  [
    "G","B","V","G","R",
    "B","V","G","R","R",
    "V","B","R","V","V",
    "G","R","R","V","G",
    "B","B","B","G","G"
  ],
  
  "goals": 
  [
    {
      "goalType": "G",
      "goalAmount": 10
    },
    {
      "goalType": "B",
      "goalAmount": 10
    },
    {
      "goalType": "V",
      "goalAmount": 10
    },
    {
      "goalType": "R",
      "goalAmount": 10
    }
  ],
  
  "startingMoveCount": 20
  
}`

func getLevel(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, levelString)
	fmt.Println("sent the test level string")
}

func main() {
	fmt.Println("Running the server...")

	http.HandleFunc("/level", getLevel)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

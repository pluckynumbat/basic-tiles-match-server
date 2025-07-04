package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

var levelStrings = []string{
	`{
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
  
}`,
	`{
  "name": "From The Server 2",
  
  "seed": 20,
  
  "colorCount": 4,

  "colorPalette": ["G","B","Y","V"],
  
  "gridLength": 7,
  
  "isStartingGridFixed": false,
  
  "goals": 
  [
    {
      "goalType": "G",
      "goalAmount": 20
    },
    {
      "goalType": "A",
      "goalAmount": 20
    }
  ],
  
  "startingMoveCount": 30
}`,
	`{
  "name": "From The Server 3",
  
  "seed": 30,
  
  "colorCount": 5,

  "colorPalette": ["G","B","Y","V","R"],
  
  "gridLength": 9,
  
  "isStartingGridFixed": false,
  
  "goals": 
  [
    {
      "goalType": "R",
      "goalAmount": 30
    },
    {
      "goalType": "Y",
      "goalAmount": 30
    }
  ],
  
  "startingMoveCount": 40
}`,
}

func getLevel(w http.ResponseWriter, req *http.Request) {
	index := rand.IntN(len(levelStrings))
	fmt.Fprintf(w, levelStrings[index])
	fmt.Println("sent the test level string")
}

func testPing(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
	fmt.Println("sent the OK response for the test ping")
}

func main() {
	fmt.Println("Running the server...")

	http.HandleFunc("/ping", testPing)
	http.HandleFunc("/level", getLevel)

	log.Fatal(http.ListenAndServe(":8090", nil))
}

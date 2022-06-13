package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type movie struct {
	Title    string
	Genre    string
	Actors   string
	Director string
	Rated    string
	Type     string
	Language string
	Ratings  string
}

//This is our request API key
const api_key string = "d3c9a85e"

func main() {
	result := GetRequest("i", "tt3896198")

	//WILL BE REMOVED LATER ON THIS IS JUST A TEST TO SEE IF WE PROPERLY GATHERED ALL INFORMATION (check terminal output when running)
	var firstMovie movie
	json.Unmarshal([]byte(result), &firstMovie)
	fmt.Println("Results: ")
	fmt.Println(firstMovie.Title)
	fmt.Println(firstMovie.Genre)
	fmt.Println(firstMovie.Actors)
	fmt.Println(firstMovie.Director)
	fmt.Println(firstMovie.Rated)
	fmt.Println(firstMovie.Type)
	fmt.Println(firstMovie.Language)
	fmt.Println(firstMovie.Ratings)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

// example request i=tt3896198
// GetRequest takes in the query code and a flag. Returns a string of JSON
func GetRequest(flag string, query string) (json string) {
	theURL := "http://www.omdbapi.com/?" + flag + "=" + query + "&apikey=" + api_key

	response, err := http.Get(theURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Printf("[STATUS CODE]: {%d}", response.StatusCode)
	fmt.Printf("[CONTENT LENGHT]: {%d}", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

	return string(content)
}

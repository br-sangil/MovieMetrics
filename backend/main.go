package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const api_key string = "d3c9a85e"

func main() {
	GetRequest("tt3896198", "i")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

// i=tt3896198
// GetRequest takes in the query code and a flag. Returns a string of JSON
func GetRequest(query string, flag string) (json string) {
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

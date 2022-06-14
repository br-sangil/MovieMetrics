package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Movie structure containing info. used for determining "rank"
type Movie struct { // Ranking System: 1/36 for a "point"

	Title    string // if match: importance(1 / 36) * 8
	Genre    string // if match: importance(1 / 36) * 7
	Actors   string // if match: importance(1 / 36) * 6
	Director string // if match: importance(1 / 36) * 5
	Rated    string // if match: importance(1 / 36) * 4
	Type     string // if match: importance(1 / 36) * 3
	Language string // if match: importance(1 / 36) * 2
	Ratings  string // if match: importance(1 / 36) * 1

	value    string //the short title for the movie used to simplify programming
	priority int    // The priority of the movie or the "score"
	index    int    //index of the movie in the PriorityQueue
}

//This is our request API key
const api_key string = "d3c9a85e"

//THE CODE BELOW MAY NEED TO BE MOVED TO A DIFFERENT MODULE
//---------------------------------------------------------------------------
type PriorityQueue []*Movie

func (h PriorityQueue) Len() int { return len(h) }

//In Order to implement the heap.Interface we must use the less func
//but we want the opposite result so we will use greater than instead
func (h PriorityQueue) Less(idxP1, idxP2 int) bool { return h[idxP1].priority > h[idxP2].priority }

//swap the given indices
func (h PriorityQueue) Swap(idxP1, idxP2 int) {
	h[idxP1], h[idxP2] = h[idxP2], h[idxP1]
	h[idxP1].index = idxP1
	h[idxP2].index = idxP2
}

//push the next movie item in the priority queue using the given h
//x is any "object" which implements heap.PriorityQueue
func (h *PriorityQueue) Push(x any) {
	n := len(*h)
	movie := x.(*Movie)
	movie.index = n
	*h = append(*h, movie)
}

//pop the next item in the priority queue using the given pq
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	movie := old[n-1]
	old[n-1] = nil
	movie.index = -1
	*pq = old[0 : n-1]
	return *movie
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(m *Movie, value string, priority int) {
	m.value = value
	m.priority = priority
	heap.Fix(pq, m.index)
}

//---------------------------------------------------------------------------

func main() {
	result := GetRequest("i", "tt3896198")

	//WILL BE REMOVED LATER ON THIS IS JUST A TEST TO SEE IF WE PROPERLY GATHERED ALL INFORMATION (check terminal output when running)
	var firstMovie Movie
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

	firstMovie.priority = 400
	movies := map[string]int{
		firstMovie.Title: firstMovie.priority,
		"movie":          550,go
		"movie2":         50,
	}

	pq := make(PriorityQueue, len(movies))
	i := 0
	for value, priority := range movies {
		pq[i] = &Movie{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	//insert new item:
	movie := &Movie{
		value:    "MOVIE",
		priority: 300,
	}
	heap.Push(&pq, movie)
	pq.update(movie, movie.value, 300)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(Movie)
		fmt.Printf("%2d:%s\n", item.priority, item.value)
	}
	//--------------------------------------^ Will be removed soon (functioning heap)-------------------------------------

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

// GetRequest takes in the query code and a flag. Returns a string of JSON
func GetRequest(flag string, query string) string {
	theURL := "http://www.omdbapi.com/?" + flag + "=" + query + "&apikey=" + api_key

	response, err := http.Get(theURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Printf("[STATUS CODE]: {%d}\n", response.StatusCode)
	fmt.Printf("[CONTENT LENGHT]: {%d}\n", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

	return string(content)
}

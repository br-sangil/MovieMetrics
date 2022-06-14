package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
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
	result := GetRequest("i=tt3896198")

	wordMap := getCommonWords()

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
		"movie":          550,
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

	getTitlePoints(movie, &firstMovie, wordMap)
	//--------------------------------------^ Will be removed soon (functioning heap)-------------------------------------

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/random", getRandomMovie)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

// Posts a random movie to the /random page
func getRandomMovie(w http.ResponseWriter, r *http.Request) {
	movieFound := false

	for !movieFound {
		imbdID := strconv.Itoa(rand.Intn(2155529) + 1)
		for len(imbdID) < 7 {
			imbdID = "0" + imbdID
		}
		// fmt.Println("Movie ID:", imbdID)
		content := GetRequest("i=tt" + imbdID)

		var data map[string]interface{}
		err := json.Unmarshal([]byte(content), &data)
		checkNilErr(err)
		// fmt.Printf("Movie data: %+v\n", data)

		if data["Response"] != "False" {
			movieFound = true

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}

// GetRequest takes in the query code and a flag. Returns a string of JSON
//flagAndQuery param must take the form as flag=query
func GetRequest(flagAndQuery string) string {
	//there could be a more effective way of concatenating the strings
	//but this is the easiest way to do it for now
	theURL := "http://www.omdbapi.com/?" + flagAndQuery + "&apikey=" + api_key

	response, err := http.Get(theURL)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Printf("[STATUS CODE]: {%d}\n", response.StatusCode)
	// fmt.Printf("[CONTENT LENGHT]: {%d}\n", response.ContentLength)
	// fmt.Println("-----output-----")
	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	// fmt.Println(string(content))
	return string(content)
}

// Terminates the application if the error is not nil
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

//gives a priority to a Movie based on the desired Movie
func getPriority(m *Movie, desiredMovie *Movie, common map[string]string) float64 {
	var priority float64
	priority += getTitlePoints(m, desiredMovie, common)
	priority += getGenrePoints(m, desiredMovie)
	priority += getActorPoints(m, desiredMovie)
	priority += getDirectorPoints(m, desiredMovie)
	priority += getRatedPoints(m, desiredMovie)
	priority += getTypePoints(m, desiredMovie)
	priority += getLanguagePoints(m, desiredMovie)
	return priority
}

// Calculates the points earned for Movie m based on the important matching words in the title
func getTitlePoints(m *Movie, desiredMovie *Movie, common map[string]string) float64 {
	//first identify common words for both m and desiredMovie
	movieNewTitle := removeCommonWords(m, common)
	desiredMovieNewTitle := removeCommonWords(desiredMovie, common)
	//then compare lexicographically to get score and return and save into m
	// (1/36 * 8) / len(desiredMovieNewTitle)
	var onePoint float64 = (1. / 36.) * 8. / float64(len(desiredMovieNewTitle))

	//if match then += onePoint
	points := 0.0
	for _, str := range movieNewTitle {
		for _, str2 := range desiredMovieNewTitle {
			if strings.ToLower(string(str)) == strings.ToLower(string(str2)) {
				points += onePoint
			}
		}
	}

	return points
}

// Removes the common words for a specified movie.
// We use our map[string]string as a set for "membership" test (does m.Title have words in common)
// used in getTitlePoints
func removeCommonWords(m *Movie, common map[string]string) string {
	movieNewTitle := ""
	for _, str := range m.Title {
		if _, ok := common[string(str)]; !ok {
			//word is not common
			movieNewTitle += string(str) + " "
		}
	}

	return strings.Trim(movieNewTitle, " ")
}

// Calculates the total of points gained for a Movie m based on the Genre
func getGenrePoints(m *Movie, desiredMovie *Movie) float64 {
	desiredGenre := strings.Split(desiredMovie.Genre, ", ")
	newGenre := strings.Split(m.Genre, ", ")

	var onePoint float64 = (1. / 36.) * 7. / float64(len(desiredGenre))

	points := 0.0
	for _, str := range desiredGenre {
		for _, str2 := range newGenre {
			if str == str2 {
				points += onePoint
			}
		}
	}

	return points
}

// Calculates the total of points gained for a Movie m based on the Actors
func getActorPoints(m *Movie, desiredMovie *Movie) float64 {
	desiredActors := strings.Split(desiredMovie.Actors, ", ")
	newActors := strings.Split(m.Actors, ", ")

	var onePoint float64 = (1. / 36.) * 6. / float64(len(desiredActors))
	points := 0.0
	for _, str := range desiredActors {
		for _, str2 := range newActors {
			if str == str2 {
				points += onePoint
			}
		}
	}

	return points
}

// Calculates the total of points gained for a Movie m based on the Director(s)
func getDirectorPoints(m *Movie, desiredMovie *Movie) float64 {
	desiredDirector := strings.Split(desiredMovie.Director, ", ")
	newDirector := strings.Split(m.Director, ", ")

	var onePoint float64 = (1. / 36.) * 5. / float64(len(desiredDirector))
	points := 0.0
	for _, str := range desiredDirector {
		for _, str2 := range newDirector {
			if str == str2 {
				points += onePoint
			}
		}
	}

	return points
}

// Calculates the total of points gained for a Movie m for what it's Rated
// if the rating matches then we will give full points
// else we give half points because who cares about the rating. RIGHT?? do you care?
func getRatedPoints(m *Movie, desiredMovie *Movie) float64 {
	var onePoint float64 = (1. / 36.) * 4
	if m.Rated == desiredMovie.Rated {
		return onePoint
	}
	return onePoint / 2.
}

// Calculates the total of points gained for a Movie m for the type of media
// if media type matches then we give full points
// else we give half points becuase we still want to consider series, episodes, etc..
func getTypePoints(m *Movie, desiredMovie *Movie) float64 {
	var onePoint float64 = (1. / 36.) * 3
	if m.Type == desiredMovie.Type {
		return onePoint
	}

	return onePoint / 2.
}

// Calculates the total of points gained for a Movie m for the language
// if m matches desiredMovie we give full points
// else do not add any points
func getLanguagePoints(m *Movie, desiredMovie *Movie) float64 {
	var onePoint float64 = (1. / 36.) * 3
	if m.Language == desiredMovie.Language {
		return onePoint
	}

	return 0
}

// fetches HTML data from espressoenglish.net for the 100 most common words in english
// we store this data (PARSED) with key value pair for efficient look up (we will use the map mainly as a set abstract data type)
func getCommonWords() map[string]string {
	resp, err := soup.Get("https://www.espressoenglish.net/the-100-most-common-words-in-english/")
	checkNilErr(err)

	// fmt.Printf("%v type: %T \n", resp, resp)
	doc := soup.HTMLParse(resp)
	words := doc.FindAll("td")

	wordMap := make(map[string]string)
	for _, word := range words {
		commonWord := strings.Split(word.Text(), ".")
		wordMap[strings.Trim(commonWord[1], " ")] = strings.Trim(commonWord[1], " ")
	}
	return wordMap
}

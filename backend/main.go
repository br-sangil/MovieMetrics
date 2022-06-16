package main

import (
	"container/heap"
	"encoding/json"
	"errors"
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
	Actors   string // if match: importance(1 / 36) * 11
	Director string // if match: importance(1 / 36) * 8
	// Rated    string // if match: importance(1 / 36) * 0 (REMOVE THIS?)
	// Type     string // if match: importance(1 / 36) * 3 (REMOVE THIS?)
	Language string // if match: importance(1 / 36) * 2
	Poster   string // the url for posters
	// Ratings  string // if match: importance(1 / 36) * 1

	// value    string  //the short title for the movie used to simplify programming (usually will be set to a shorter version of the title)
	priority float64 // The priority of the movie or the "score"
	index    int     //index of the movie in the PriorityQueue
	// isAdd    bool    // if added already is true
}

// structure which defines the movies searched via API
type MovieSearched struct {
	Search []*Movie `json:"Search"`
}

// structure which defines the movies that pertain to an actor via a different API
type ActorSearch struct {
	Results []*Actor `json:"results"`
}

// structure used specifically for The Movie Database API since we have different Keys
type Actor struct {
	Name   string    `json:"name"`
	Movies []*Movie2 `json:"known_for"`
}

type Movie2 struct {
	Title string `json:"original_title"`
}

//This is our request API key
const api_key string = "d3c9a85e"

// This is our request API key for The Movie Database API
const api_key2 string = "007058439d1b582e951c080688c60593"

//Port number being used
const port string = ":8081"

//THE CODE BELOW MAY NEED TO BE MOVED TO A DIFFERENT MODULE
//---------------------------------------------------------------------------
type PriorityQueue []*Movie

func (h PriorityQueue) Len() int { return len(h) }

//In Order to implement the heap.Interface we must use the less func
//but we want the opposite result so we will use greater than instead
func (h PriorityQueue) Less(idxP1, idxP2 int) bool {
	return float32(h[idxP1].priority) > float32(h[idxP2].priority)
}

// func (h PriorityQueue) Less(i, j int) bool {
// 	res := big.NewFloat(h[i].priority).Cmp(big.NewFloat(h[j].priority))
// 	if res == 0 {
// 		return false
// 	} else if res > 0 {
// 		return true
// 	}
// 	return false
// }

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
	return movie
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(m *Movie, title string, priority float64) {
	m.Title = title
	m.priority = priority
	heap.Fix(pq, m.index)
}

//---------------------------------------------------------------------------

func main() {
	wordMap := getCommonWords()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r != nil {
			defer r.Body.Close()
			getMovieSearch(w, r, wordMap)
		}
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
		if r != nil {
			defer r.Body.Close()
		}

	})

	http.HandleFunc("/random", getRandomMovie)

	log.Fatal(http.ListenAndServe(port, nil))
}

// search up movie /?t=<whatever you're searching up by title>
func getMovieSearch(w http.ResponseWriter, r *http.Request, common map[string]string) {
	enableCors(&w)

	if r != nil {
		query, ok := r.URL.Query()["t"]

		if !ok || len(query[0]) < 1 {
			fmt.Println("QUERY KEY t NOT FOUND IN URL")
			return
		}

		searchTerm := query[0]
		fmt.Println("The title search term is", searchTerm)

		content := GetRequest("t=" + searchTerm)

		var desiredMovie Movie
		err := json.Unmarshal([]byte(content), &desiredMovie)
		checkNilErr(err)
		fmt.Printf("Movie data: %+v\n", desiredMovie)

		searchKeys := removeCommonWords(&desiredMovie, common)
		fmt.Println("Search terms:", searchKeys)

		// use search keys to search for movies
		// such that we search for movies with all keywords, then one keyword at a time
		movies, err := searchMovies(searchKeys)
		checkNilErr(err)

		for _, word := range strings.Split(searchKeys, " ") {
			searchResults, err := searchMovies(word)
			checkNilErr(err)

			movies = append(movies, searchResults[:]...)
		}

		actors := strings.Split(desiredMovie.Actors, ", ")
		for _, actor := range actors {
			actorMovies := findMoviesByActor(actor).Results[0]
			for _, movie := range actorMovies.Movies {
				content := GetRequest("t=" + movie.Title)

				var actorMovie Movie
				err = json.Unmarshal([]byte(content), &actorMovie)
				checkNilErr(err)
				fmt.Printf("%+v\n", actorMovie)

				movies = append(movies, &actorMovie)
			}
		}

		// this part is for priority queue
		movieMap := make(map[string]*Movie)
		for _, movie := range movies {
			if _, ok := movieMap[movie.Title]; !ok {
				movieMap[movie.Title] = movie
			}
		}

		pq := make(PriorityQueue, 0)
		heap.Init(&pq)

		for _, movie := range movieMap {
			getPriority(movie, &desiredMovie, common)
			// add movie to priority queue here
			heap.Push(&pq, movie)
			pq.update(movie, movie.Title, movie.priority)
		}
		heap.Init(&pq)

		for i := 0; i < pq.Len(); i++ {
			println("movie:", pq[i].Title, "pr:", pq[i].priority)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(pq)
	}
}

// GET request themoviedb.org, searches by actor
func findMoviesByActor(actor string) ActorSearch {
	words := strings.Split(actor, " ")
	actorParse := strings.Join(words, "%20")
	println(actorParse)

	dbURL := "https://api.themoviedb.org/3/search/person?api_key=" + api_key2 + "&language=en-US&query=" + actorParse + "&page=1&include_adult=false&region=US"
	response, err := http.Get(dbURL)
	checkNilErr(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	var resultActors ActorSearch
	json.Unmarshal(content, &resultActors)

	return resultActors
}

// Posts a random movie to the /random page (magnificient use of imbdID)
func getRandomMovie(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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
	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	return string(content)
}

// Searches and returns a list of movies based on key words
func searchMovies(keyWord string) ([]*Movie, error) {
	result := []byte(GetRequest("s=" + keyWord))
	if result == nil || string(result) == "" {
		return nil, errors.New("BAD REQUEST")
	}

	var movies MovieSearched
	json.Unmarshal(result, &movies)

	for _, val := range movies.Search {
		response := GetRequest("t=" + val.Title)
		json.Unmarshal([]byte(response), &val)
	}
	return movies.Search, nil
}

//gives a priority to a Movie based on the desired Movie
func getPriority(m *Movie, desiredMovie *Movie, common map[string]string) {
	var priority float64
	// Title    string // if match: importance(1 / 36) * 8 (2)
	// Genre    string // if match: importance(1 / 36) * 7 (3)
	// Actors   string // if match: importance(1 / 36) * 11 (1)
	// Director string // if match: importance(1 / 36) * 8 (3.5)
	priority += getTitlePoints(m, desiredMovie, common)
	priority += getGenrePoints(m, desiredMovie)
	priority += getActorPoints(m, desiredMovie)
	priority += getDirectorPoints(m, desiredMovie)
	// priority += getRatedPoints(m, desiredMovie)
	// priority += getTypePoints(m, desiredMovie)
	priority += getLanguagePoints(m, desiredMovie)
	(*m).priority = priority * 100
	// println("calculated priority ", priority)
}

// Calculates the points earned for Movie m based on the important matching words in the title
func getTitlePoints(m *Movie, desiredMovie *Movie, common map[string]string) float64 {
	//first identify common words for both m and desiredMovie
	movieNewTitle := removeCommonWords(m, common)
	desiredMovieNewTitle := removeCommonWords(desiredMovie, common)

	var onePoint float64 = ((1. / 36.) * 8.) / float64(len(desiredMovieNewTitle))

	//if match then += onePoint
	points := 0.0
	for _, str := range movieNewTitle {
		for _, str2 := range desiredMovieNewTitle {
			if strings.EqualFold(string(str), string(str2)) {
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
	for _, str := range strings.Split(m.Title, " ") {
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

	var onePoint float64 = ((1. / 36.) * 7.) / float64(len(desiredGenre))

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

	var onePoint float64 = ((1. / 36.) * 11.) / float64(len(desiredActors))
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

	var onePoint float64 = ((1. / 36.) * 8.) / float64(len(desiredDirector))
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
// func getRatedPoints(m *Movie, desiredMovie *Movie) float64 {
// 	var onePoint float64 = (1. / 36.) * 4
// 	if m.Rated == desiredMovie.Rated {
// 		return onePoint
// 	}
// 	return onePoint / 2.
// }

// Calculates the total of points gained for a Movie m for the type of media
// if media type matches then we give full points
// else we give half points becuase we still want to consider series, episodes, etc..
// func getTypePoints(m *Movie, desiredMovie *Movie) float64 {
// 	var onePoint float64 = (1. / 36.) * 4
// 	if m.Type == desiredMovie.Type {
// 		return onePoint
// 	}

// 	return onePoint / 2.
// }

// Calculates the total of points gained for a Movie m for the language
// if m matches desiredMovie we give full points
// else do not add any points
func getLanguagePoints(m *Movie, desiredMovie *Movie) float64 {
	var onePoint float64 = (1. / 36.) * 2
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

	doc := soup.HTMLParse(resp)
	words := doc.FindAll("td")

	wordMap := make(map[string]string)
	for _, word := range words {
		commonWord := strings.Split(word.Text(), ".")
		wordMap[strings.Trim(commonWord[1], " ")] = strings.Trim(commonWord[1], " ")
	}
	return wordMap
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

// Terminates the application if the error is not nil
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

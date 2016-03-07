package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
	if err != nil {
		fmt.Println("Error while loading site")
	}
	defer resp.Body.Close()
	fmt.Println("status code is ", resp.Status)

	movie := new(Movies)
	getJSON("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json", movie)
	fmt.Printf("The movie: %v was released in %v - the IMDB rating is %v%% with %v votes \n", movie.Title, movie.Year, movie.ImdbRating*10, movie.ImdbVotes)
}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

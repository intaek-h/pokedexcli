package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationArea struct {
	Count    int
	Next     string
	Previous string
	Results  []LocationAreaResult
}

type LocationAreaResult struct {
	Name string
	URL  string
}

var nextUrl = "https://pokeapi.co/api/v2/location-area"
var prevUrl = "https://pokeapi.co/api/v2/location-area"

func GetMaps() {
	res, err := http.Get(nextUrl)

	if err != nil {
		log.Fatal(err)
		return
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("HTTP Error: %s", body)
		return
	}
	if err != nil {
		log.Fatal(err)
		return
	}

	data := []byte(body)
	result := LocationArea{}
	err = json.Unmarshal(data, &result)

	if err != nil {
		log.Fatal(err)
		return
	}

	totalCount := int(result.Count)

	fmt.Printf("Location Areas(%d/%d):\n", totalCount, len(result.Results))
	fmt.Println("==============")

	for _, area := range result.Results {
		fmt.Println(area.Name)
	}

	fmt.Println("==============")

	nextUrl = result.Next

	if len(result.Previous) > 0 {
		prevUrl = result.Previous
	}

	fmt.Println("Next: ", result.Next)
	fmt.Println("Prev: ", result.Previous)
}

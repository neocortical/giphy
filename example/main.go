package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dattch/giphy"
)

func main() {
	var giphyClient = giphy.NewClient(os.Getenv("GIPHY_API_KEY"))

	var gifID = "srmSEu2ZtOq64"
	gif, err := giphyClient.GIF(gifID, nil)
	if err != nil {
		log.Printf("error getting gif '%s' by ID: %v", gifID, err)
	} else {
		log.Printf("got gif response:\n%s", dumpObj(gif))
	}

	var searchQuery = "blep"
	search, err := giphyClient.Search(searchQuery, &giphy.Options{
		Limit: 3,
	})
	if err != nil {
		log.Printf("error searching for '%s' gifs: %v", searchQuery, err)
	} else {
		log.Printf("got search response:\n%s", dumpObj(search))
	}

	trending, err := giphyClient.Trending(&giphy.Options{
		Limit: 3,
	})
	if err != nil {
		log.Printf("error getting trending gifs: %v", err)
	} else {
		log.Printf("got trending response:\n%s", dumpObj(trending))
	}

	random, err := giphyClient.Random([]string{"dog"}, &giphy.Options{
		Limit: 3,
	})
	if err != nil {
		log.Printf("error getting random gifs: %v", err)
	} else {
		log.Printf("got random response:\n%s", dumpObj(random))
	}

	translate, err := giphyClient.Translate("perro", &giphy.Options{
		Limit: 3,
	})
	if err != nil {
		log.Printf("error getting translate gifs: %v", err)
	} else {
		log.Printf("got translate response:\n%s", dumpObj(translate))
	}
}

func dumpObj(i interface{}) string {
	data, _ := json.MarshalIndent(i, "", "  ")
	return string(data)
}

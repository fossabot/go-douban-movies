package model

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL = "https://movie.douban.com/top250"

type Movie struct {
	Title        string
	OtherNames   []string
	Director     string
	Starring     string
	Year         int
	Areas        []string
	Tags         []string
	Star         float64
	CommentCount int
	Quote        string
}

var Movies = make([]Movie, 0, 250)

func InitMovies() {
	pages := pageURLs(baseURL)
	for _, page := range pages {
		url := baseURL + page
		r, err := http.Get(url)
		if err != nil {
			log.Fatalf("http get err: %v", err)
		}
		doc, err := goquery.NewDocumentFromReader(r.Body)
		if err != nil {
			log.Fatalf("new document err: %v", err)
		}
		Movies = append(Movies, movies(doc)...)
	}
}

package model

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL = "https://movie.douban.com/top250"

type Movie struct {
	Title     string
	Subtitle  string
	OtherName string
	Desc      string
	Year      string
	Area      string
	Tag       string
	Star      string
	Comment   string
	Quote     string
}

var Movies = make([]Movie, 0, 250)

func Init() {
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

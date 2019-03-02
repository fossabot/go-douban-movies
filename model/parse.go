package model

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// 返回那 10 个页面的 URL
func pageURLs(url string) []string {
	r, err := http.Get(url)
	if err != nil {
		log.Fatalf("http get err: %v", err)
	}
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Fatalf("new document err: %v", err)
	}

	urls := make([]string, 0, 10)
	urls = append(urls, "")
	doc.Find("#content > div > div.article > div.paginator > a").
		Each(func(i int, s *goquery.Selection) {
			href, ok := s.Attr("href")
			if !ok {
				log.Fatalf("href: %v isn't exist", href)
			}
			urls = append(urls, href)
		})
	return urls
}

// 返回某个页面的 25 个 Movie 结构
func movies(doc *goquery.Document) []Movie {
	movies := make([]Movie, 0, 25)
	doc.Find("#content > div > div.article > ol > li").
		Each(func(i int, s *goquery.Selection) {
			title := s.Find(".hd a span").Eq(0).Text()

			subtitle := s.Find(".hd a span").Eq(1).Text()
			subtitle = strings.TrimLeft(subtitle, " / ")

			otherName := s.Find(".hd a span").Eq(2).Text()
			otherName = strings.TrimLeft(otherName, "  / ")

			desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
			DescInfo := strings.Split(desc, "\n")
			desc = DescInfo[0]
			movieDesc := strings.Split(DescInfo[1], "/")
			year := strings.TrimSpace(movieDesc[0])
			area := strings.TrimSpace(movieDesc[1])
			tag := strings.TrimSpace(movieDesc[2])

			star := s.Find(".bd .star .rating_num").Text()

			comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
			cmp := regexp.MustCompile("[0-9]")
			comment = strings.Join(cmp.FindAllString(comment, -1), "")

			quote := s.Find(".quote .inq").Text()

			movie := Movie{
				Title:     title,
				Subtitle:  subtitle,
				OtherName: otherName,
				Desc:      desc,
				Year:      year,
				Area:      area,
				Tag:       tag,
				Star:      star,
				Comment:   comment,
				Quote:     quote,
			}
			movies = append(movies, movie)
		})
	return movies
}

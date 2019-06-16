package model

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// 返回 10 个页面的 URL
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
	urls = append(urls, "") // 当前页面
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

// 码点为 160，而不是正常的 32（空格）
const magicSpace = " "

// 返回某个页面的 25 个 Movie 结构
func movies(doc *goquery.Document) []Movie {
	movies := make([]Movie, 0, 25)
	doc.Find("#content > div > div.article > ol > li").
		Each(func(i int, s *goquery.Selection) {
			// 肖申克的救赎
			title := s.Find(".hd a span").Eq(0).Text()

			// &nbsp;/&nbsp;The Shawshank Redemption&nbsp;/&nbsp;月黑高飞(港)  /  刺激1995(台)
			otherNamesMeta := s.Find(".hd a span").Eq(1).Text() + s.Find(".hd a span").Eq(2).Text()
			otherNamesMeta = replaceSpace(otherNamesMeta)
			otherNamesMeta = strings.TrimPrefix(otherNamesMeta, "/ ")
			otherNames := strings.Split(otherNamesMeta, "/")
			trimAll(otherNames)

			// 导演: 弗兰克·德拉邦特 Frank Darabont&nbsp;&nbsp;&nbsp;主演: 蒂姆·罗宾斯 Tim Robbins /...<br>
			//                             1994&nbsp;/&nbsp;美国&nbsp;/&nbsp;犯罪 剧情
			descMeta := s.Find(".bd p").Eq(0).Text()
			descMeta = strings.TrimSpace(descMeta)
			descMetas := strings.Split(descMeta, "\n")

			// 导演: 弗兰克·德拉邦特 Frank Darabont&nbsp;&nbsp;&nbsp;主演: 蒂姆·罗宾斯 Tim Robbins /...
			dsInfo := strings.Split(descMetas[0], magicSpace+magicSpace+magicSpace)
			director := dsInfo[0]
			director = strings.TrimPrefix(director, "导演: ")
			starring := ""
			if len(dsInfo) >= 2 {
				starring = dsInfo[1]
				starring = strings.TrimPrefix(starring, "主演: ")
				starring = strings.Replace(starring, " /", "", 1)
			}

			// 1994&nbsp;/&nbsp;美国&nbsp;/&nbsp;犯罪 剧情
			yatInfo := strings.Split(strings.TrimSpace(descMetas[1]), "/")
			yearStr := strings.Trim(yatInfo[0], magicSpace)
			year, _ := strconv.Atoi(yearStr)
			areas := strings.Split(yatInfo[1], " ")
			trimAll(areas)
			tags := strings.Split(yatInfo[2], " ")
			trimAll(tags)

			// 9.6
			starStr := s.Find(".bd .star .rating_num").Text()
			star, err := strconv.ParseFloat(starStr, 64)
			if err != nil {
				log.Fatalf("%s ParseFloat err: %v", starStr, err)
			}

			// 1451552人评价
			commentCountStr := s.Find(".bd .star span").Eq(3).Text()
			commentCountStr = strings.TrimSuffix(commentCountStr, "人评价")
			commentCount, err := strconv.Atoi(commentCountStr)
			if err != nil {
				log.Fatalf("%s Atoi err: %v", commentCountStr, err)
			}

			// 希望让人自由。
			quote := s.Find(".quote .inq").Text()

			movie := Movie{
				Title:        title,
				OtherNames:   otherNames,
				Director:     director,
				Starring:     starring,
				Year:         year,
				Areas:        areas,
				Tags:         tags,
				Star:         star,
				CommentCount: commentCount,
				Quote:        quote,
			}
			movies = append(movies, movie)
		})
	return movies
}

func trimAll(strs []string) {
	for i := range strs {
		strs[i] = strings.Trim(strs[i], " ")
	}
	for i := range strs {
		strs[i] = strings.Trim(strs[i], magicSpace)
	}
}

func replaceSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

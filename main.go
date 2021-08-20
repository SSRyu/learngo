package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var (
	query   string = "python"
	limit   int    = 50
	baseURL string = fmt.Sprintf("https://jp.indeed.com/jobs?q=%s&limit=%d&", query, limit)
)

func main() {
	// fmt.Println(baseURL)
	totalPages := getPages()
	fmt.Println(totalPages)
	// for i := 0; i < totalPages; i++ {
	for i := 0; i < 1; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := fmt.Sprint(baseURL, "start=", limit*page)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		title := strings.TrimSpace(card.Find(".title>a").Text())
		location := strings.TrimSpace(card.Find(".sjcl").Text())
		fmt.Println(i, id, title, location)
	})
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	pageRegexp := regexp.MustCompile(`[0-9,]+`)
	doc.Find("#searchCountPages").Each(func(i int, s *goquery.Selection) {
		fmt.Println()
		for _, str := range strings.Split(strings.TrimSpace(s.Text()), " ") {
			matched := pageRegexp.MatchString(str)
			if matched {
				num, _ := strconv.Atoi(strings.ReplaceAll(str, ",", ""))
				pages = (num / 50) + 1
				break
			}
		}
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.Status)
	}
}

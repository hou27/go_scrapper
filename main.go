package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type requestResult struct {
	url    string
	status string
}

var baseURL string = "https://kr.indeed.com/jobs?q=go&limit=50"
var errRequestFailed = errors.New("Request failed")

func main() {
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.itoa(page*50) // integer to ascii
	fmt.Println("Requesting", pageURL)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL) // func http.Get(url string) (resp *http.Response, err error)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // when this func finished, close it. Prevent memory leaks.

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination-list").Each(func(idx int, sel *goquery.Selection){
		// fmt.Println(sel.Html())
		pages = sel.Find("a").Length()
	}) // func (*goquery.Selection).Each(f func(int, *goquery.Selection)) *goquery.Selection

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
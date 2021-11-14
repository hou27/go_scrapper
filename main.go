package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	link     string
	title    string
	location string
	salary   string
	summary  string
}

type requestResult struct {
	url    string
	status string
}

var baseURL string = "https://kr.indeed.com/jobs?q=go&limit=50"
var errRequestFailed = errors.New("Request failed")

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...) // extractedjobs... ::: the values inside of extractedjobs
	}
	
	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob

	pageURL := baseURL + "&start=" + strconv.Itoa(page*50) // integer to ascii
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")

	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
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

func extractJob(card *goquery.Selection) extractedJob {
	link, _ := card.Attr("href")
		title:= card.Find(".jobTitle>span").Text()
		location := card.Find(".companyLocation").Text()
		salary := cleanString(card.Find(".salary-snippet").Text())
		summary := card.Find(".job-snippet").Text()
	return extractedJob{
		link:     link,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() // Flush writes any buffered data to the underlying io.Writer.

	headers := []string{"LINK", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com" + job.link, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func cleanString(str string) string {
	// clean the space from the sides, clean the gap between strings(return arr), concatenates the elements of its first argument to create a single string.
	/**
	FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c) and returns an array of slices of s. If all code points in s satisfy f(c) or the string is empty, an empty slice is returned.
	 */
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
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
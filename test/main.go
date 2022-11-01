package main

import (
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/shadowfish07/go-readability"
)

func main() {
	// res, err := http.Get("https://www.yinfans.me/movie/35761")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// if res.StatusCode != 200 {
	// 	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	// }

	res, _ := os.Open("test.html")

	// Load the HTML document
	doc, _ := goquery.NewDocumentFromReader(res)
	r := readability.New(doc, nil)
	r.IsProbablyReadable()
}

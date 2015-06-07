package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os" // for command-line arguments
)

/* we will need
* A main function, which reads parameters from the CLI
* A function which parses the page for links
* A function which visits links
 */

func main() {
	domain := os.Args[1]
	fmt.Println("Parsing", domain)
	page, err := goquery.NewDocument(domain)

	if err != nil {
		panic(err)
	}

	ch := make(chan string)
	parse(ch, page)
}

func parse(ch chan<- string, page *goquery.Document) {
	page.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Println("link found:", href)
		}
	})
}

package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os" // for command-line arguments
	"runtime"
)

/* we will need
* A main function, which reads parameters from the CLI
* A function which parses the page for links
* A function which visits links
 */

func main() {
	setMaxProc()

	domain := os.Args[1]
	fmt.Println("Parsing", domain)
	page, err := goquery.NewDocument(domain)

	if err != nil {
		panic(err)
	}

	link_q := make(chan string, 100)
	parse(link_q, page)

	for n := 0; n < 100; n++ {
		go func(jobs <-chan string) {
			url := <-jobs

			fmt.Printf("Worker %v popped %v\n", n, url)
		}(link_q)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("Exiting")
}

func parse(jobs chan<- string, page *goquery.Document) {
	page.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Println("link found:", href)
			jobs <- href
		}
	})
}

func setMaxProc() {
	num_cpu := runtime.NumCPU()
	fmt.Println("Setting GOMAXPROCS to", num_cpu)
	runtime.GOMAXPROCS(num_cpu)
}

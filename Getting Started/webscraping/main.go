package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:name`
	Rating string `json:rating`
	Price  string `json:price`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("amazon.com"),
	)
	// fmt.Println("Here")

	c.OnHTML("div[class=p13n-sc-truncate-desktop-type2  p13n-sc-truncated]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("https://www.amazon.com/books-used-books-textbooks/b?ie=UTF8&node=283155")
}

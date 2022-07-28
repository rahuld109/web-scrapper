package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Date   string `json:"date"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("dev.to"),
	)

	var items []item

	c.OnHTML("div.crayons-story__body", func(h *colly.HTMLElement) {
		item := item{
			Title:  h.ChildText("div.crayons-story__indention h2.crayons-story__title a[href]"),
			Author: h.ChildText("div.profile-preview-card button[id]"),
			Date:   h.ChildText("a[href] time[datetime]"),
		}

		items = append(items, item)

	})

	c.Visit("https://dev.to/")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err)
	}

	os.WriteFile("scapper.json", content, 0644)
}

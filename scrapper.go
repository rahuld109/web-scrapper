package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type item struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Date        string   `json:"date"`
	ProfileUrl  string   `json:"profile_url"`
	UserPageUrl string   `json:"user_page_url"`
	Tags        []string `json:"tags"`
	PageUrl     string   `json:"page_url"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("dev.to"),
	)

	var items []item

	c.OnHTML("div.crayons-story__body", func(h *colly.HTMLElement) {

		item := item{
			Title:       h.ChildText("div.crayons-story__indention h2.crayons-story__title a[href]"),
			Author:      h.ChildText("div.profile-preview-card button[id]"),
			ProfileUrl:  h.ChildAttr("a.crayons-avatar img", "src"),
			UserPageUrl: h.Request.AbsoluteURL(h.ChildAttr("a.crayons-avatar", "href")),
			Date:        h.ChildAttr("time", "datetime"),
			Tags:        strings.Split(h.ChildText("div.crayons-story__tags a.crayons-tag"), "#")[1:],
			PageUrl:     h.Request.AbsoluteURL(h.ChildAttr("div.crayons-story__indention h2.crayons-story__title a", "href")),
		}

		items = append(items, item)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://dev.to/top/week")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err)
	}

	os.WriteFile("output.json", content, 0644)
}

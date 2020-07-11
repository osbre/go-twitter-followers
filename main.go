package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

const usernameQuery = "table.user-item > tbody > tr:nth-child(2) > td.info > a > span.username"

func main() {
	var users []string

	c := colly.NewCollector(
		colly.AllowedDomains("mobile.twitter.com"),
		colly.Async(true),
	)
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML(usernameQuery, func(e *colly.HTMLElement) {
		users = append(users, strings.Trim(e.Text, "@"))
	})

	c.OnHTML("div.w-button-more > a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"));
	})

	usernames := os.Args[1:]
	for _, username := range usernames {
		_ = c.Visit("https://mobile.twitter.com/" + username + "/followers")
	}

	c.Wait()

	fmt.Print(users)
}

package main

import (
	"fmt"
	"gosearch/dbwrapper"

	"github.com/gocolly/colly"
)

func main() {

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("[Main] Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		title := r.URL.Host
		link := r.URL.String()
		dbwrapper.CreateWebLink(title, link)
		fmt.Println("[Main] Visiting", title, link)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")

	/*
		rows, err := db.Query("SELECT id, first_name FROM users LIMIT $1", 3)
		defer rows.Close()

		for rows.Next() {
			var id int
			var firstName string
			err = rows.Scan(&id, &firstName)
			if err != nil {
				// handle this error
				panic(err)
			}
			fmt.Println(id, firstName)
		}

		err = rows.Err()
		if err != nil {
			panic(err)
		}
	*/

}

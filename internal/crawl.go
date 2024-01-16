package internal

import (
    "log"

    "github.com/gocolly/colly/v2"
)

func scrapeURLs(baseURL string, amount int) []string {
    c := colly.NewCollector()

    var urls []string

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        urls = append(urls, link)
        if len(urls) >= amount {
            c.DisableAllowedDomains()
        }
    })

    c.Visit(baseURL)

    if len(urls) > amount {
        urls = urls[:amount]
    }

    return urls
}
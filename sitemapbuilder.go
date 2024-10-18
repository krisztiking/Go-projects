package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func sitemapbuilder() {
	urlFlag := flag.String("url", "https://gophercises.com", "The url that you want to build a sitemap for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	fmt.Println("Request URl: ", reqUrl.String())
	fmt.Println("Base URL: ", base)

	links, _ := Parse(resp.Body)
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}

}

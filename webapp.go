package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	//Titles    []string `xml:"url>news>title"`
	//Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}
type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var results Sitemapindex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(bytes, &results)
	// newsMap := make(map[string]News)

	for _, Location := range results.Locations {
		fmt.Println(Location)
		// resp, _ := http.Get(Location)
		resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/politics.xml")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Body)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		fmt.Println(n)
		//resp.Body.Close()
	}
	//resp.Body.Close()

}

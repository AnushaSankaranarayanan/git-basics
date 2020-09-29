package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}
type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
`)

func main() {
		i := 0
	   	for i < 10 {
	   		fmt.Println(i)
	   		i++
	   	}
	   	x := 5
	   	for {
	   		fmt.Println("Do stuff", x)
	   		x += 3
	   		if x > 25 {
	   			break
	   		}
	   	}
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//bytes := washPostXML
	var results Sitemapindex
	xml.Unmarshal(bytes, &results)
	for _, location := range results.Locations {
		fmt.Println(location.Loc)
	}
	//fmt.Println(results.Locations)
	resp.Body.Close()

}

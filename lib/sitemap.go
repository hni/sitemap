package sitemap

import (
	"encoding/xml"
	"log"
)

type Url struct {
	XMLName    xml.Name `xml:"url"`
	Loc        string   `xml:"loc"`
	LastMod    string   `xml:"lastmod"`
	ChangeFreq string   `xml:"changefreq"`
	Priority   float32  `xml:"priority"`
}

type UrlSet struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	Urls    []Url    `xml:"url"`
}

func Sitemap(urls []string)([]byte) {
	u := &UrlSet{}
    for _, url := range urls {
		u.Urls = append(u.Urls, Url{Loc: url})
	}
	output, err := xml.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

    return output
}

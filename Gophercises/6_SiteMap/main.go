package main

import (
	"Sitemap/Link"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// <?xml version="1.0" encoding="UTF-8"?>
// <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
// 		<url>
//   		<loc>http://www.example.com/</loc>
// 		</url>
// 		<url>
// 		  	<loc>http://www.example.com/dogs</loc>
// 		</url>
// </urlset>

type CustomUrl struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type Urlset struct {
	Xmlns   string   `xml:"xmlns,attr"`
	XMLName xml.Name `xml:"urlset"`
	Loc     []CustomUrl
}

func main() {
	rootWebsiteLink := flag.String("Root", "https://www.calhoun.io/", "Root Website To Start Tracking")
	depth := flag.Int("Depth", 2, "Number of websites in stack")
	flag.Parse()

	mp := make(map[string]bool)
	mp[*rootWebsiteLink] = true
	bfs(*rootWebsiteLink, mp, *depth)

	fmt.Println(mp)
	urls := make([]CustomUrl, len(mp))
	iter := 0
	for key := range mp {
		urls[iter] = CustomUrl{Loc: key}
		iter++
	}
	urlset := Urlset{Loc: urls, Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}

	_, err := xml.MarshalIndent(urlset, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(xml.Header + string(data))
}

func bfs(mainLink string, visitedLinks map[string]bool, maxDepth int) {
	q := map[string]bool{
		mainLink: true,
	}
	nq := map[string]bool{}

	for i := 0; i <= maxDepth; i++ {
		for key := range q {
			buildMap(key, visitedLinks, nq)
		}
		q = nq
		nq = map[string]bool{}
	}
}

func buildMap(mainLink string, visitedLinks map[string]bool, queue map[string]bool) {

	body := giveHtml(mainLink)
	links := Link.Parser(body)

	for _, link := range links {
		if len(link.Href) == 0 {
			fmt.Println("found empty link")
			continue
		} else if link.Href[0] == '#' {
			fmt.Println("found #")
			continue
		}
		mainUrl, err := url.Parse(mainLink)
		if err != nil {
			log.Fatal(err)
		}
		childUrl, err := url.Parse(link.Href)
		if err != nil {
			log.Fatal(err)
		}
		if link.Href[0] == '/' {
			link.Href = mainUrl.Scheme + "://" + mainUrl.Host + link.Href
		} else if mainUrl.Scheme+"://"+mainUrl.Host != childUrl.Scheme+"://"+childUrl.Host {
			continue
		}

		if link.Href[len(link.Href)-1] != '/' {
			link.Href += "/"
		}

		if _, found := visitedLinks[link.Href]; !found {
			visitedLinks[link.Href] = true
			queue[link.Href] = true
		}
	}
}

func giveHtml(link string) io.Reader {
	resp, err := http.Get(link)
	if err != nil {
		defer resp.Body.Close()
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp.Body
}

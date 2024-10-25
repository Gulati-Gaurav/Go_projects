package Link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parser(reqBody io.Reader) []Link {
	doc, err := html.Parse(reqBody)
	if err != nil {
		panic(err)
	}
	links := []Link{}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.CommentNode {
			return
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			text := returnString(n)

			href := ""
			for i := range n.Attr {
				if n.Attr[i].Key == "href" {
					href = n.Attr[i].Val
					break
				}
			}

			link := Link{
				Text: text,
				Href: href,
			}
			links = append(links, link)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(doc)

	// We are getting many spaces in the links. To remove this we split the strings about consecutive spaces. go has a method for us - fields in string
	for i := range links {
		words := strings.Fields(links[i].Text)
		links[i].Text = strings.Join(words, " ")
	}
	return links
}

func returnString(n *html.Node) string {
	if n.Type == html.ElementNode {
		ans := ""
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ans += returnString(c)
		}
		return ans
	} else if n.Type == html.TextNode {
		return n.Data
	}
	return ""
}

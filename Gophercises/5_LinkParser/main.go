package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	htmlStr := `<html>
	<head>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
	</head>
	<body>
	<h1>Social stuffs</h1>
	<div>
		<a href="https://www.twitter.com/joncalhoun">
		Check me out on twitter
		<i class="fa fa-twitter" aria-hidden="true"></i>
		</a>
		<a href="https://github.com/gophercises">
		Gophercises is on <strong>Github</strong>!
		</a>
	</div>
	</body>
	</html>`

	reqBody := strings.NewReader(htmlStr)
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
			// TODO - instead of the picking the first node, iterate
			href := n.Attr[0].Val
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
	fmt.Printf("%+v", links)
}

func returnString(n *html.Node) string {
	if n.Type == html.CommentNode {
		return ""
	} else if n.Type == html.ElementNode {
		ans := ""
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ans += returnString(c)
		}
		return ans
	} else if n.Type == html.TextNode {
		return n.Data // todo
	}
	return ""
}

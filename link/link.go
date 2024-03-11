package link

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// Returns an array of Link structs
func ParseURLs(r io.Reader) ([]Link, error) {
	// Parse the HTML
	z := html.NewTokenizer(r)

	// An array of link structs
	links := make([]Link, 0)

	depth := 0
	index := -1

	// Tokenize the HTML and print out the text content
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links, errors.New("error token encountered")
		case html.CommentToken:
			fmt.Println("Comment: ", string(z.Text()))
		case html.TextToken:
			if depth > 0 {
				links[index].Text = string(z.Text())
				depth = depth - 1
			}
		case html.StartTagToken, html.EndTagToken:
			for _, attr := range z.Token().Attr {
				if attr.Key == "href" && len(attr.Val) > 0 && attr.Val[0] != '#' {
					index = len(links)
					depth = depth + 1
					links = append(links, Link{attr.Val, "text"})
				}
			}
		}
	}
}

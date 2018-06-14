package functions

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

// https://godoc.org/golang.org/x/net/html
func ParseHTML(body io.ReadCloser) *html.Node {
	root, err := html.Parse(body)
	if err != nil {
		// Skriv ut mappnamnet också
		log.Println(err.Error())
	}

	return root
}

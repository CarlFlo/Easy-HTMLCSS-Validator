package functions

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// ParseHTML parses the html
func ParseHTML(body io.Reader) {

	doc, err := html.Parse(body)
	if err != nil {
		// ...
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" && len(n.Attr) > 0 && n.Attr[0].Val == "grouped msg_err" {

			fmt.Println(n.FirstChild)

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}

// ParseHTMLRaw will parse raw html
func ParseHTMLRaw(html string) {

	html, totErrors := filterHTML(html)

	log.Println(">", totErrors, "<")

	if len(html) == 0 {
		log.Println("ERROR! html length after parsing is 0")
		return
	}
	getGroupMsg(html)

}

func getGroupMsg(html string) {

}

// Gets section in HMTL return that contains the errors
func filterHTML(html string) (string, string) {

	atRightPoint := false
	divs := 0
	totErrorsAndWarnings := ""

	scanner := bufio.NewScanner(strings.NewReader(html)) // reads html
	html = ""                                            // resets html to "" bc we dont need it anymore
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "<div id=\"result\">") {
			atRightPoint = true
		}
		if !atRightPoint {

			// Gets total errors and warnings
			if strings.Contains(scanner.Text(), "warning(s)") {
				totErrorsAndWarnings = scanner.Text()
				cpy := totErrorsAndWarnings
				for i := 0; i < len(cpy); i++ {
					if strings.Compare(string(cpy[i]), " ") == 0 {
						totErrorsAndWarnings = totErrorsAndWarnings[1:] // Tar bort bokstaven innan på index 0 (substring)
					} else {
						break
					}
				}
			}
			continue
		}

		if strings.Contains(scanner.Text(), "<div ") {
			divs++
		} else if strings.Contains(scanner.Text(), "</div>") {
			divs--
		}

		html += fmt.Sprintf("%s\n", scanner.Text())
		if divs == 0 {
			return html, totErrorsAndWarnings
		}
	}

	return "", totErrorsAndWarnings

	/*
		li
			span
				img
			/span
			span	// has error group msg
				p
					a
				/p
				div class="ve
					p		// Onödigt
					/p
				/div
				ul
					li	// Alla induviduella fel
						em	/em
						span /spawn
					/li
					*kan finnas flera li här*
				/ul
			/li
	*/
}

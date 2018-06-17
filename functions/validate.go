package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// ValidateHTML will validate html with the server
func ValidateHTML(html string, singleHTML *HTMLVerify) error {

	resp, err := http.PostForm("https://validator.w3.org/check",
		url.Values{ // name
			"charset":  {"(detect automatically)"},
			"fragment": {html},               // HTML koden
			"doctype":  {"XHTML 1.0 Strict"}, // "HTML5" eller "XHTML 1.0 Strict"
			"group":    {"1"},                // Gruppera felen tillsammans
			"st":       {"0"},                // Clean up Markup with HTML-Tidy
			"outline":  {"1"},                // Får ut antalet h1-hx element längst ner
		})

	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Parses the html code
	ParseHTML(string(body), singleHTML)

	/* // Debug for now. This will save return html to a file*/
	name := strings.Split(singleHTML.Path, "\\")
	ioutil.WriteFile(string(fmt.Sprintf("./Done-%s-%v", RandomString(4), name[len(name)-1])), body, 0644) // debug

	return nil
}

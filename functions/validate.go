package functions

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// ValidateHTML will validate html with the server
func ValidateHTML(html, path string) error {

	resp, err := http.PostForm("https://validator.w3.org/check",
		url.Values{ // name
			"charset":  {"(detect automatically)"},
			"fragment": {html},               // HTML koden
			"doctype":  {"XHTML 1.0 Strict"}, // "HTML5" eller "XHTML 1.0 Strict"
			"group":    {"1"},                // Gruppera felen tillsammans
			"st":       {"1"},
			"outline":  {"1"}, // Får ut antalet h1-hx element längst ner
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

	ParseHTMLRaw(string(body))
	//ParseHTML(bytes.NewReader(body))

	// Debug for now. Will save return html to a file

	/*
		name := strings.Split(path, "\\")
		ioutil.WriteFile(string(fmt.Sprintf("./Done-%s-%v", RandomString(4), name[len(name)-1])), body, 0644) // debug
	*/

	return nil
}

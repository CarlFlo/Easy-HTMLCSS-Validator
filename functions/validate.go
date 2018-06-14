package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// ValidateHTML will validate html with the server
func ValidateHTML(html, path string, tmp int) error {

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

	//functions.ParseHTML(resp.Body) // body istället

	ioutil.WriteFile(string(fmt.Sprintf("./verified/%v.html", tmp)), body, 0644) // debug

	return nil
}

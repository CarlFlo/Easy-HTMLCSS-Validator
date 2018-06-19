package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

// ValidateHTMLStrict will validate html with the server for XHTML 1.0 Strict
func ValidateHTMLStrict(html string, singleHTML *HTMLVerify) error {
	/*
		// DEBUG
		ParseHTML("", singleHTML)
		runtime.Goexit()
		return nil
		// DEBUG END
	*/

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

	return nil
}

// ValidateHTML5Web will validate html with the server for XHTML 1.0 Strict
func ValidateHTML5Web(html string, singleHTML *HTMLVerify) error {

	resp, err := http.PostForm("https://validator.w3.org/check",
		url.Values{
			"fragment":        {html}, // HTML koden
			"prefill":         {"0"},
			"doctype":         {"Inline"},  // "HTML5" eller "XHTML 1.0 Strict"
			"prefill_doctype": {"html401"}, // "HTML5" eller "XHTML 1.0 Strict"
			"group":           {"1"},       // Gruppera felen tillsammans
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
	ioutil.WriteFile(string(fmt.Sprintf("./Done-html5-%s-%v", RandomString(4), name[len(name)-1])), body, 0644) // debug

	return nil
}

// Very slow and demanding
// ValidateHTML5 will validate html with jar file for html5
func ValidateHTML5(html string, singleHTML *HTMLVerify) error {

	cmd := exec.Command("java", "-jar", "./vnu.jar", "--format", "json", singleHTML.Path)

	out, err := cmd.CombinedOutput()

	if err != nil {

		// exit status 1 is also error for when jar file is not found
		if err.Error() != "exit status 1" { // status 1 is ok, means there were errors in html doc
			singleHTML.HTML5Verify.ErrorValidating = err
			return err
		}
	}

	if len(string(out)) > 16 { // Empty json result is 16 long
		singleHTML.HTML5Verify.HasWarningsOrErrors = true
	} else {
		singleHTML.HTML5Verify.HasWarningsOrErrors = false
	}
	singleHTML.HTML5Verify.Verified = true

	return nil
}

// ValidateCSS Validates css
func ValidateCSS(css string, singleHTML *HTMLVerify) error {

	return nil
}

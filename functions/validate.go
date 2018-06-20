package functions

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
)

// ValidateHTMLStrict will validate html with the server for XHTML 1.0 Strict
func ValidateHTMLStrict(html string, singleHTML *HTMLVerify) error {

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
	parseHTML(string(body), singleHTML)

	return nil
}

// validateHTML5 will validate html with the server for XHTML 1.0 Strict
func validateHTML5(html string, singleHTML *HTMLVerify) error {

	extraParams := map[string]string{
		"fragment":        html,
		"doctype":         "Inline",
		"group":           "1",
		"prefill":         "0",
		"prefill_doctype": "html401",
		"showsource":      "no",
	}

	request, err := newMultipartForm("https://validator.w3.org/nu/#textarea", extraParams)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		singleHTML.HTML5Verify.Verified = false
		return err
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			singleHTML.HTML5Verify.Verified = false
			return err
		}
		resp.Body.Close()

		// Check if it has errors here

		// singleHTML.HTML5Verify.HasWarningsOrErrors = true / false
		txtBody, _ := ioutil.ReadAll(body) // read all
		parseHTML5(txtBody, singleHTML)
	}
	singleHTML.HTML5Verify.Verified = true
	return nil
}

// validateCSS Validates css // Save return error to singleCSS.ErrorValidating variable
func validateCSS(css string, singleCSS *CSSVerify) error {
	extraParams := map[string]string{
		"text":        css,
		"profile":     "css3svg",
		"usermedium":  "all",
		"type":        "none",
		"warning":     "1",
		"vextwarning": "",
		"lang":        "en",
	}

	request, err := newMultipartForm("https://jigsaw.w3.org/css-validator/validator", extraParams)
	if err != nil {
		singleCSS.Verified = false
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		singleCSS.Verified = false
		return err
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			singleCSS.Verified = false
			return err
		}
		resp.Body.Close()

		// Read all
		txtBody, _ := ioutil.ReadAll(body)

		// Check if it has errors here
		parseCSS(txtBody, singleCSS)
	}
	singleCSS.Verified = true
	return nil
}

// Creates the form data
func newMultipartForm(uri string, params map[string]string) (*http.Request, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

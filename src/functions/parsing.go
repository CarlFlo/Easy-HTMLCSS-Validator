package functions

import (
	"bytes"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// parseHTML will parse raw html
func parseHTML(html string, singleHTML *HTMLVerify) {

	// Parses for XHTML 1.0 Strict
	parseGroupMessages([]byte(html), singleHTML)

	// Mark this html file as done
	singleHTML.AllVerified = true
}

// Will parse all the group messages
func parseGroupMessages(body []byte, singleHTML *HTMLVerify) {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	// Get Errors and warnings as long string
	doc.Find("tr .invalid").Each(func(i int, s *goquery.Selection) {

		singleHTML.StrictVerify.Result = removeWhitespace(s.Text())
		return // Only one
	})

	doc.Find("#warnings li.msg_warn p span.msg").Each(func(i int, s *goquery.Selection) {
		singleHTML.StrictVerify.Warnings = append(singleHTML.StrictVerify.Warnings, removeWhitespace(s.Text()))
	})

	doc.Find("#warnings li.msg_info p span.msg").Each(func(i int, s *goquery.Selection) {
		singleHTML.StrictVerify.Infos = append(singleHTML.StrictVerify.Infos, s.Text())
	})

	doc.Find("#error_loop li.msg_err, #error_loop li.msg_warn").Each(func(i int, s *goquery.Selection) {
		//fmt.Printf("\nGROUP: %s\n", removeWhitespace(s.Find(".msg").Text()))

		// Create ErrorType
		errorGroup := ErrorGroup{
			ErrorType: s.Find(".msg").Text(),
		}

		// Iterate over the errors
		groupInt := 0
		theError := TheError{}

		// theErrors being overwritten

		s.Find("ul li em, ul li > span, ul li > pre").Each(func(ii int, ss *goquery.Selection) {
			if groupInt > 2 {
				groupInt = 0
				theError = TheError{}
			}

			switch groupInt {
			case 0: // The line
				theError.Line = removeWhitespace(ss.Text())
			case 1: // The error
				theError.Error = removeWhitespace(ss.Text())
			case 2: // In the html
				theError.TextFromHTML = removeWhitespace(ss.Text())
				errorGroup.ErrorStrings = append(errorGroup.ErrorStrings, theError)
			default:
			}
			groupInt++
		})
		// Save group data
		singleHTML.StrictVerify.Errors = append(singleHTML.StrictVerify.Errors, errorGroup)
	})

	singleHTML.StrictVerify.Verified = true
}

func parseHTML5(html []byte, singleHTML *HTMLVerify) {

	if strings.Contains(string(html), "There were errors") {
		// Found error(s)
		singleHTML.HTML5Verify.HasWarningsOrErrors = true
	}
}

func parseCSS(css []byte, singleCSS *CSSVerify) {

	if strings.Contains(string(css), "Sorry! We found the following errors") {
		// Found error(s)
		singleCSS.HasWarningsOrErrors = true
	}
}

// Removed extra spaces from string to clean it up
func removeWhitespace(input string) string {
	reLeadcloseWhtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	reInsideWhtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := reLeadcloseWhtsp.ReplaceAllString(input, "")
	final = reInsideWhtsp.ReplaceAllString(final, " ")
	return final
}

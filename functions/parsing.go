package functions

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// ParseHTML will parse raw html
func ParseHTML(html string, singleHTML *HTMLVerify) {

	// Parses for XHTML 1.0 Strict
	totErrorsAndWarnings := filterHTML(html, singleHTML)

	/*
		if len(html) == 0 {
			log.Println("ERROR! html length after parsing is 0, path:", singleHTML.Path)
			return
		} */

	// Get total warnings and errors as said from the HTML Doc
	singleHTML.StrictVerify.TotalErrors, singleHTML.StrictVerify.TotalWarnings = parseErrAndWarn(totErrorsAndWarnings)

	//log.Println("Err:", singleHTML.TotalErrors, " Warn:", singleHTML.TotalWarnings)

	// Mark this html file as done
	singleHTML.AllVerified = true
}

// Seperated the errors from the warnings and returns them as int
func parseErrAndWarn(str string) (int, int) {
	errVal := 0
	var errBuff bytes.Buffer
	errDone := false
	warnVal := 0
	warnTurn := false
	var warnBuff bytes.Buffer

	// str format -> xx Errors, yy warning(s)

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' { // Not a space
			if !errDone {
				errBuff.WriteByte(str[i])
				continue
			} else if !warnTurn && str[i] == ',' { // Warnings
				warnTurn = true
				continue
			} else if warnTurn {
				warnBuff.WriteByte(str[i])
			}
		} else { // a space
			if !errDone {
				tmpErrVal, err := strconv.Atoi(errBuff.String())
				if err != nil {
					errVal = -1
				} else {
					errVal = tmpErrVal
				}
				errDone = true
			} else if warnTurn && len(warnBuff.String()) != 0 {
				tmpWarnVal, err := strconv.Atoi(warnBuff.String())
				if err != nil {
					warnVal = -1
				} else {
					warnVal = tmpWarnVal
				}
				return errVal, warnVal
			}
		}
	}
	return errVal, warnVal
}

// Gets section in HMTL return that contains the errors
func filterHTML(html string, singleHTML *HTMLVerify) string {

	atErrors := false
	atWarnings := false
	divs := 0
	totErrorsAndWarnings := ""

	scanner := bufio.NewScanner(strings.NewReader(html)) // reads html
	html = ""                                            // resets html to "" bc we dont need it anymore
	// Used to keep track on what index the warn and err is supposed to be writed to in the singleHTML object
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "<div id=\"result\">") {
			// At section where errors are
			atErrors = true
			continue
		} else if strings.Contains(scanner.Text(), "<h3 id=\"preparse_warnings\">Notes and Potential Issues</h3>") {
			// At section where warnings are
			atWarnings = true
			continue
		}

		if atWarnings {
			// Notes and Potential Issues - Section in html response
			// Parse warning data here

			// Take entire row and filter out tags and unwanted text
			if strings.Contains(scanner.Text(), "<span class=\"msg\">") {

				// Cleanup string. Removed /span and /p from string
				trimmedStr := scanner.Text()[:len(scanner.Text())-11]

				// Goes thru string backwards
				// Looking for  "msg">  string
				for i := len(scanner.Text()) - 3; i > 0; i-- {
					if scanner.Text()[i] == '>' && scanner.Text()[i-1] == '"' && scanner.Text()[i-2] == 'g' && scanner.Text()[i-3] == 's' {
						trimmedStr = trimmedStr[i+1:]
						break
					}
				}

				singleHTML.StrictVerify.Warnings = append(singleHTML.StrictVerify.Warnings, trimmedStr)

				// Check before loop
				if strings.Contains(scanner.Text(), "</span>") {
					continue
				}
				panic("Warning did not have </span> at the end") // debug
				/*
					// Loop till closing span
					for scanner.Scan() {
						//singleHTML.Warnings[currentWarningIndex] += scanner.Text()
						if strings.Contains(scanner.Text(), "</span>") { // end
							break
						}
					}
					continue
				*/
			}

			// Detects end of warnings
			if strings.Contains(scanner.Text(), "<!-- End of \"warnings\". -->") {
				atWarnings = false
				continue
			}
			continue
		}

		if atErrors {
			// Parse error data here

			if strings.Contains(scanner.Text(), "<li class=\"grouped msg_err\">") {

				for scanner.Scan() {
					if strings.Contains(scanner.Text(), "<span class=\"msg\">") {
						// Error section. save this scanner.Text()
						// Replace Errors with new struct type
					}
				}

				// Iterate to end off li. Each error is in its own li inside a ul
			}

			// Checks for when error section is over
			if strings.Contains(scanner.Text(), "<div ") {
				divs++
			} else if strings.Contains(scanner.Text(), "</div>") {
				divs--
			}
			html += fmt.Sprintf("%s\n", scanner.Text())
			if divs == 0 {
				return totErrorsAndWarnings
			}
			continue
		}

		// This is done outside of warnings and errors //

		// Check for <br> elements

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

	// XHTML 1.0 Strict verify is done
	singleHTML.StrictVerify.Verified = true

	return totErrorsAndWarnings

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

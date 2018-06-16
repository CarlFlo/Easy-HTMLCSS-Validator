package functions

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// ParseHTML will parse raw html
func ParseHTML(html string, singleHTML *HTMLVerify) {

	html, totErrorsAndWarnings := filterHTML(html)

	if len(html) == 0 {
		log.Println("ERROR! html length after parsing is 0, path:", singleHTML.Path)
		return
	}

	// Get total warnings and errors as said from the HTML Doc
	singleHTML.TotalErrors, singleHTML.TotalWarnings = parseErrAndWarn(totErrorsAndWarnings)

	//log.Println("Err:", singleHTML.TotalErrors, " Warn:", singleHTML.TotalWarnings)

	getGroupMsg(html)

	// Mark as done
	singleHTML.Verified = true
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

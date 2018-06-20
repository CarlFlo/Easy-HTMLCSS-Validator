(function () {
  'use strict';

  test()

  function test() {
    let css = document.getElementById("css")
    let html5 = document.getElementById("html5")
    let htmlStrict = document.getElementById("htmlStrict")

    document.getElementById("folderName").innerHTML = OUTPUT_RESULT.Projects[0].FolderName

    // CSS
    if (OUTPUT_RESULT.Projects[0].CSSs[0] != null) {
      if (OUTPUT_RESULT.Projects[0].CSSs[0].Verified) {
        if (OUTPUT_RESULT.Projects[0].CSSs[0].HasWarningsOrErrors) {
          css.getElementsByTagName("h3")[0].innerHTML = "File has errors or warnings"
          css.getElementsByTagName("p")[0].innerHTML = `Path ${OUTPUT_RESULT.Projects[0].CSSs[0].Path}`
          css.setAttribute("class", "hasProblems")
        } else {
          css.getElementsByTagName("h3")[0].innerHTML = "Validated OK!"
          css.getElementsByTagName("p")[0].innerHTML = `Path ${OUTPUT_RESULT.Projects[0].CSSs[0].Path}`
          css.setAttribute("class", "noProblems")
        }
      } else {
        css.getElementsByTagName("h3")[0].innerHTML = "Could not validate file"
        css.getElementsByTagName("p")[0].innerHTML = `Path ${OUTPUT_RESULT.Projects[0].CSSs[0].Path}`
      }
    } else {
      css.getElementsByTagName("h3")[0].innerHTML = "No CSS file"
    }

    // HTML5
    if (OUTPUT_RESULT.Projects[0].HTMLs[0].HTML5Verify.Verified) {
      if (OUTPUT_RESULT.Projects[0].HTMLs[0].HTML5Verify.HasWarningsOrErrors) {
        html5.getElementsByTagName("h3")[0].innerHTML = "File has errors or warnings"
        html5.getElementsByTagName("p")[0].innerHTML = `Path ${OUTPUT_RESULT.Projects[0].HTMLs[0].Path}`
        html5.setAttribute("class", "hasProblems")
      } else {
        html5.getElementsByTagName("h3")[0].innerHTML = "Validated OK!"
        html5.getElementsByTagName("p")[0].innerHTML = `Path ${OUTPUT_RESULT.Projects[0].HTMLs[0].Path}`
        html5.setAttribute("class", "noProblems")
      }
    } else {
      html5.getElementsByTagName("h3")[0].innerHTML = "Could not validate file"
      html5.getElementsByTagName("p")[0].innerHTML = `Path ${OUTPUT_RESULT.Projects[0].HTMLs[0].Path}`
    }

    // XHTML 1.0 Strict
    if (OUTPUT_RESULT.Projects[0].HTMLs[0].StrictVerify.Verified) {

      htmlStrict.getElementsByTagName("h3")[0].innerHTML = "Result:"
      htmlStrict.getElementsByTagName("p")[0].innerHTML = `${OUTPUT_RESULT.Projects[0].HTMLs[0].StrictVerify.Result}`
    } else {
      htmlStrict.getElementsByTagName("h2")[0].innerHTML = "Could not validate file"
      htmlStrict.getElementsByTagName("p")[0].innerHTML = `Path ${OUTPUT_RESULT.Projects[0].HTMLs[0].Path}`
    }
  }

}());
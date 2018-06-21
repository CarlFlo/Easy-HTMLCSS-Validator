(function () {
  'use strict';

  let CONTAINER = document.getElementById("container");

  displayResult();
  handleExpandCollapse();

  function handleExpandCollapse() {

    // Add eventlisteners to every "button" for expand and collapse
    let list = document.getElementsByClassName("expandCollapse"); // expandCollapse
    for (let i = 0; i < list.length; i++) {
      list[i].addEventListener("click", function() {
        this.classList.toggle("active");
        let content = this.nextElementSibling;
        if (content.style.display === "block") {
          content.style.display = "none";
        } else {
          content.style.display = "block";
        }
      });
    }
  }

  function displayResult() {

    // Iterate over all projects
    for (let i = 0; i < OUTPUT_RESULT.Projects.length; i++) {

      let miniContainer = document.createElement("div");
      miniContainer.setAttribute("id", `projectId-${i}`);
      miniContainer.setAttribute("class", `project`);

      // Sets foldername as title for the project section
      let title = document.createElement("h1");
      title.setAttribute("class", "folderName");
      title.innerHTML = OUTPUT_RESULT.Projects[i].FolderName;
      miniContainer.appendChild(title);

      // [CSS] - Check if css exists
      if (OUTPUT_RESULT.Projects[i].CSSs.length > 0) {
        let cssList = document.createElement("ul"); // list with everything css
        cssList.setAttribute("class", "cssList");

        let cssSection = document.createElement("h2");
        cssSection.setAttribute("class", "fileGroup");
        cssSection.innerHTML = "[CSS]";
        cssList.appendChild(cssSection);

        // Iterate thru array
        for (let j = 0; j < OUTPUT_RESULT.Projects[i].CSSs.length; j++) {
          let tmpCSS = document.createElement("li");
          let cssFile = document.createElement("h3");

          if (!OUTPUT_RESULT.Projects[i].CSSs[j].Verified) {
            cssFile.innerHTML = "[VALIDATE FAILED]: "
          }

          cssFile.innerHTML += `${OUTPUT_RESULT.Projects[i].CSSs[j].Path}`

          if (!OUTPUT_RESULT.Projects[i].CSSs[j].Verified || OUTPUT_RESULT.Projects[i].CSSs[j].HasWarningsOrErrors) {
            cssFile.setAttribute("class", "hasProblems");
          } else {
            cssFile.setAttribute("class", "noProblems");
          }

          tmpCSS.appendChild(cssFile);
          cssList.appendChild(tmpCSS);
        }
        miniContainer.appendChild(cssList);  // Append list
      } else {
        // no css. Show no css msg on screen
      }

      // [HTML] - Check if html files exists
      if (OUTPUT_RESULT.Projects[i].HTMLs.length > 0) {
        let htmlList = document.createElement("ul");  // htmlList - list with everything html
        miniContainer.appendChild(htmlList);
        htmlList.setAttribute("class", "htmlList");

        let fileGroup = document.createElement("h2"); // fileGroup
        htmlList.appendChild(fileGroup);
        fileGroup.setAttribute("class", "fileGroup");
        fileGroup.innerHTML = "[HTML]";

        let htmlContainer = document.createElement("ul"); // htmlContainer
        htmlList.appendChild(htmlContainer);
        htmlContainer.setAttribute("class", "htmlContainer");

        // For every html document in folder
        for (let j = 0; j < OUTPUT_RESULT.Projects[i].HTMLs.length; j++) {

          let singleHTML = document.createElement("li");  // singleHTML
          htmlContainer.appendChild(singleHTML);
          singleHTML.setAttribute("class", "singleHTML");

          let fileName = document.createElement("h3");  // fileName
          singleHTML.appendChild(fileName);
          fileName.setAttribute("class", "fileName expandCollapse");
          fileName.innerHTML = `Path: ${OUTPUT_RESULT.Projects[i].HTMLs[j].Path}`;

          let verifyList = document.createElement("ul");  // verifyList
          singleHTML.appendChild(verifyList);
          verifyList.setAttribute("class", "verifyList");

          // For html5 and XHTML 1.0 Strict
          let verifyHTML5 = document.createElement("li"); // verifyHTML5
          verifyList.appendChild(verifyHTML5);
          verifyHTML5.setAttribute("class", "verifyHTML5")

          // Show result
          let html5ValidateResult = document.createElement("h4");
          verifyHTML5.appendChild(html5ValidateResult);
          html5ValidateResult.innerHTML = "HTML5 Result: ";

          // HTML5
          if (OUTPUT_RESULT.Projects[i].HTMLs[j].HTML5Verify.Verified) {

            if (OUTPUT_RESULT.Projects[i].HTMLs[j].HTML5Verify.HasWarningsOrErrors) {
              // Validated with errors
              html5ValidateResult.innerHTML += "Validated with errors or warnings!";
              html5ValidateResult.setAttribute("class", "hasProblems");
            } else {
              //Validated witout errors
              html5ValidateResult.innerHTML += "OK!";
              html5ValidateResult.setAttribute("class", "noProblems");
            }
          } else {
            // HTML5 verify failed
            html5ValidateResult.innerHTML += "Couldn't validate file!";
            html5ValidateResult.setAttribute("class", "hasProblems");
          }

          // XHTML 1.0 Strict
          let verifyXHTMLstrict = document.createElement("li"); // verifyXHTMLstrict
          verifyList.appendChild(verifyXHTMLstrict);
          verifyXHTMLstrict.setAttribute("class", "verifyXHTMLstrict")

          // How result
          let XHTMLresult = document.createElement("h4");
          verifyXHTMLstrict.appendChild(XHTMLresult);
          XHTMLresult.innerHTML = `XHTML 1.0 Strict Result: `;

          let tmpResultText = document.createElement("div");
          tmpResultText.innerHTML = `${OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Result}`;
          tmpResultText.setAttribute("class", "hasProblems");
          XHTMLresult.appendChild(tmpResultText);

          // List that holds all warnings, info and errors
          let errAndWarnList = document.createElement("ul");  // errAndWarnList
          verifyXHTMLstrict.appendChild(errAndWarnList);
          errAndWarnList.setAttribute("class", "errAndWarnList");

          // Populate list

          // Warnings
          let xhtmlWarningList = document.createElement("li");  // xhtmlWarningList
          errAndWarnList.appendChild(xhtmlWarningList);
          xhtmlWarningList.setAttribute("class", "xhtmlWarningList");

          let warningsText = document.createElement("h4");
          xhtmlWarningList.appendChild(warningsText);
          warningsText.setAttribute("class", "xhtml-warning");
          warningsText.innerHTML = "Warning(s):";

          if (OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Warnings.length == 0) {
            // No warnings
            let tmp = document.createElement("p");
            xhtmlWarningList.appendChild(tmp);
            tmp.innerHTML = "No warnings"
            tmp.setAttribute("class", "noProblems");

          } else {
            let tmpWarnsList = document.createElement("ul");
            xhtmlWarningList.appendChild(tmpWarnsList);
            for (let k = 0; k < OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Warnings.length; k++) {
              let tmpWarnsLi = document.createElement("li");
              tmpWarnsList.appendChild(tmpWarnsLi);
              let tmpWarnsP = document.createElement("p");
              tmpWarnsLi.appendChild(tmpWarnsP);
              tmpWarnsP.innerHTML = OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Warnings[k];
            }
          }

          // Info
          let xhtmlInfoList = document.createElement("li"); // xhtmlInfoList
          errAndWarnList.appendChild(xhtmlInfoList);
          xhtmlInfoList.setAttribute("class", "xhtmlInfoList");

          let infoText = document.createElement("h4");
          xhtmlInfoList.appendChild(infoText);
          infoText.setAttribute("class", "xhtml-info");
          infoText.innerHTML = "Info(s):";

          if (OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Infos.length == 0) {
            // No infos
            let tmp = document.createElement("p");
            xhtmlInfoList.appendChild(tmp);
            tmp.innerHTML = "No info";
            tmp.setAttribute("class", "noProblems");

          } else {
            let tmpInfoList = document.createElement("ul");
            xhtmlInfoList.appendChild(tmpInfoList);
            for (let k = 0; k < OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Infos.length; k++) {
              let tmpInfoLi = document.createElement("li");
              tmpInfoList.appendChild(tmpInfoLi);
              let tmpInfoP = document.createElement("p");
              tmpInfoLi.appendChild(tmpInfoP);
              tmpInfoP.innerHTML = OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Infos[k];
            }
          }

          // Errors
          let xhtmlErrorList = document.createElement("li");  // xhtmlErrorList
          errAndWarnList.appendChild(xhtmlErrorList);
          xhtmlErrorList.setAttribute("class", "xhtmlErrorList");

          let errorText = document.createElement("h4"); // xhtml-error
          xhtmlErrorList.appendChild(errorText);
          errorText.setAttribute("class", "xhtml-error expandCollapse");
          errorText.innerHTML = "Error(s):";

          if (OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors.length == 0) {
            // no errors
            let tmp = document.createElement("p");
            xhtmlErrorList.appendChild(tmp);
            tmp.innerHTML = "No errors";
            tmp.setAttribute("class", "noProblems");

          } else {

            let tmpErrorList = document.createElement("ul");  // holds all errors
            xhtmlErrorList.appendChild(tmpErrorList);

            for (let k = 0; k < OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors.length; k++) {
              let tmpErrLi = document.createElement("li");
              tmpErrorList.appendChild(tmpErrLi);
              let groupName = document.createElement("h5");
              tmpErrLi.appendChild(groupName);
              groupName.innerHTML = OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors[k].ErrorType;
              groupName.setAttribute("class", "errGroup");

              let tmpInduvidualErrors = document.createElement("ul");
              tmpErrLi.appendChild(tmpInduvidualErrors);

              for (let kk = 0; kk < OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors[k].ErrorStrings.length; kk++) {
                let tmpErrInfo = document.createElement("li");
                tmpInduvidualErrors.appendChild(tmpErrInfo);

                let tmpLine = document.createElement("p");
                tmpErrInfo.appendChild(tmpLine);
                let tmpError = document.createElement("p");
                tmpErrInfo.appendChild(tmpError);
                let tmpText = document.createElement("p");
                tmpErrInfo.appendChild(tmpText);

                tmpLine.innerText = `Line: ${OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors[k].ErrorStrings[kk].Line}`;
                tmpError.innerText = `Error: ${OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors[k].ErrorStrings[kk].Error}`;
                tmpText.innerText = `Text: ${OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors[k].ErrorStrings[kk].TextFromHTML}`;
                //tmpText.
              } 
            }
          }

        }
        miniContainer.appendChild(htmlList); // Append list
      } else {
        // no html. Show no html msg on screen
      }


      /* appendar till container */
      CONTAINER.appendChild(miniContainer);
    }
  }

}());
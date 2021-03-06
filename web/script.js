(function () {
  'use strict';

  let CONTAINER = document.getElementById("container");

  displayResult();
  handleExpandCollapse();

  function handleExpandCollapse() {

    // Add eventlisteners to every "button" for expand and collapse
    let list = document.getElementsByClassName("expandCollapse"); // expandCollapse
    for (let i = 0; i < list.length; i++) {
      list[i].classList.add("isExpanded");

      // Removed need for first time doubleclick to enable the feature on each button
      let tmp = list[i].nextElementSibling;

      if (tmp == undefined) {
        continue;
      }

      tmp.style.display = "block";

      list[i].addEventListener("click", function () {
        this.classList.toggle("active");

        let content = this.nextElementSibling;
        if (content.style.display === "block") {
          content.style.display = "none";
          this.classList.add("isCollapsed");
          this.classList.remove("isExpanded");
        } else {
          content.style.display = "block";
          this.classList.remove("isCollapsed");
          this.classList.add("isExpanded");
        }
      });
    }
  }

  function displayResult() {



    if (typeof OUTPUT_RESULT == 'undefined') {
      console.log("Couldn't load ../output.js");

      let errMsg = document.createElement("h1");
      CONTAINER.appendChild(errMsg);
      errMsg.setAttribute("align", "center");
      errMsg.classList.add("hasProblems");
      errMsg.classList.add("blink_me");
      errMsg.innerText = "Couldn't load ../output.js";

      return
    }

    // Iterate over all projects
    for (let i = 0; i < OUTPUT_RESULT.Projects.length; i++) {

      let miniContainer = document.createElement("div");
      miniContainer.setAttribute("id", `projectId-${i}`);
      miniContainer.setAttribute("class", `project`);

      /* append to container */
      CONTAINER.appendChild(miniContainer);

      // So entire thing can be collapsed and expanded
      //let stuffContainer = document.createElement("div");
      //miniContainer.appendChild(stuffContainer)

      // Sets foldername as title for the project section
      let title = document.createElement("h1");
      title.setAttribute("class", "folderName expandCollapse");
      title.innerHTML = OUTPUT_RESULT.Projects[i].FolderName;
      miniContainer.appendChild(title);

      let cssHtmlHolder = document.createElement("div");
      miniContainer.appendChild(cssHtmlHolder)

      // [CSS] - Check if css exists
      if (OUTPUT_RESULT.Projects[i].CSSs.length > 0) {
        let cssList = document.createElement("ul"); // list with everything css
        cssHtmlHolder.appendChild(cssList);  // Append list
        cssList.setAttribute("class", "cssList");

        let cssSection = document.createElement("h2");
        cssSection.setAttribute("class", "fileGroup expandCollapse");
        cssSection.innerHTML = "[CSS]";
        cssList.appendChild(cssSection);

        // So collapse and extend will work
        let cssDIV = document.createElement("div");
        cssList.appendChild(cssDIV);

        // Iterate thru array
        for (let j = 0; j < OUTPUT_RESULT.Projects[i].CSSs.length; j++) {
          let tmpCSS = document.createElement("li");
          let cssFile = document.createElement("h3");

          tmpCSS.appendChild(cssFile);
          cssDIV.appendChild(tmpCSS);

          if (!OUTPUT_RESULT.Projects[i].CSSs[j].Verified) {
            cssFile.innerHTML = "[VALIDATION FAILED]: "
          }

          cssFile.innerHTML += `${OUTPUT_RESULT.Projects[i].CSSs[j].Path}`

          if (!OUTPUT_RESULT.Projects[i].CSSs[j].Verified || OUTPUT_RESULT.Projects[i].CSSs[j].HasWarningsOrErrors) {
            cssFile.setAttribute("class", "hasProblems");
          } else {
            cssFile.setAttribute("class", "noProblems");
          }


        }

      } else {
        // no css. Show no css msg on screen
      }

      // [HTML] - Check if html files exists
      if (OUTPUT_RESULT.Projects[i].HTMLs.length > 0) {
        let htmlList = document.createElement("ul");  // htmlList - list with everything html
        cssHtmlHolder.appendChild(htmlList); // Append list
        htmlList.setAttribute("class", "htmlList");

        let fileGroup = document.createElement("h2"); // fileGroup
        htmlList.appendChild(fileGroup);
        fileGroup.setAttribute("class", "fileGroup expandCollapse");
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
              html5ValidateResult.classList.add("blink_me");
            } else {
              //Validated witout errors
              html5ValidateResult.innerHTML += "OK!";
              html5ValidateResult.setAttribute("class", "noProblems");
            }
          } else {
            // HTML5 verify failed
            html5ValidateResult.innerHTML += "Couldn't validate file!";
            html5ValidateResult.setAttribute("class", "hasProblems");
            html5ValidateResult.classList.add("blink_me");
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

          // List that holds all warnings, info and errors (and outline)
          let errAndWarnList = document.createElement("ul");  // errAndWarnList
          verifyXHTMLstrict.appendChild(errAndWarnList);
          errAndWarnList.setAttribute("class", "errAndWarnList");

          // Populate list

          // Outline
          let outlineList = document.createElement("li");  // outlineList
          errAndWarnList.appendChild(outlineList);
          outlineList.setAttribute("class", "outlineList");

          let outlineTest = document.createElement("h4");
          outlineList.appendChild(outlineTest);
          outlineTest.setAttribute("class", "expandCollapse");
          outlineTest.innerHTML = "Outline:";

          if (OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Outline != undefined) {
            if (OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Outline.length == 0) {
              // Nothing in outline
              console.log("outline length is 0")
              outlineTest.innerHTML += " len 0"
            } else {

              let tmpOutlineList = document.createElement("ul");
              outlineList.appendChild(tmpOutlineList);

              
              let tmpOutlineLi = document.createElement("li");
              tmpOutlineList.appendChild(tmpOutlineLi);
              let tmpOutlineP = document.createElement("p");
              tmpOutlineLi.appendChild(tmpOutlineP);
              tmpOutlineP.innerText = OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Outline;
            }
          } else {
            // outline error
            console.log("outline error")
            outlineTest.innerHTML += " undefined"
          }


          // Warnings
          let xhtmlWarningList = document.createElement("li");  // xhtmlWarningList
          errAndWarnList.appendChild(xhtmlWarningList);
          xhtmlWarningList.setAttribute("class", "xhtmlWarningList");

          let warningsText = document.createElement("h4");
          xhtmlWarningList.appendChild(warningsText);
          warningsText.setAttribute("class", "xhtml-warning expandCollapse");
          warningsText.innerHTML = "Warning(s):";

          if (OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Warnings != undefined) {
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
                tmpWarnsP.innerText = OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Warnings[k];
              }
            }
          } else {
            // No warnings
            console.log("no warnings")
            let errMsg = document.createElement("p");
            xhtmlWarningList.appendChild(errMsg);
            errMsg.innerText = "No warnings!";
          }

          // Info
          let xhtmlInfoList = document.createElement("li"); // xhtmlInfoList
          errAndWarnList.appendChild(xhtmlInfoList);
          xhtmlInfoList.setAttribute("class", "xhtmlInfoList");

          let infoText = document.createElement("h4");
          xhtmlInfoList.appendChild(infoText);
          infoText.setAttribute("class", "xhtml-info expandCollapse");
          infoText.innerHTML = "Info(s):";

          if (OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Infos != undefined) {
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
                tmpInfoP.innerText = OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Infos[k];
              }
            }
          } else {
            // no infos
            console.log("no info")
            let errMsg = document.createElement("p");
            xhtmlInfoList.appendChild(errMsg);
            errMsg.innerText = "No info!";
          }

          // Errors
          let xhtmlErrorList = document.createElement("li");  // xhtmlErrorList
          errAndWarnList.appendChild(xhtmlErrorList);
          xhtmlErrorList.setAttribute("class", "xhtmlErrorList");

          let errorText = document.createElement("h4"); // xhtml-error
          xhtmlErrorList.appendChild(errorText);
          errorText.setAttribute("class", "xhtml-error expandCollapse");
          errorText.innerHTML = "Error(s):";

          if (OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors != undefined) {

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
                groupName.setAttribute("class", "errGroup expandCollapse");
  
                let tmpInduvidualErrors = document.createElement("ul");
                tmpErrLi.appendChild(tmpInduvidualErrors);
  
                for (let kk = 0; kk < OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors[k].ErrorStrings.length; kk++) {
                  let tmpErrInfo = document.createElement("li");
                  tmpInduvidualErrors.appendChild(tmpErrInfo);
  
                  let tmpError = document.createElement("p");
                  tmpErrInfo.appendChild(tmpError);
                  tmpError.classList.add("markMe")
                  let tmpText = document.createElement("p");
                  tmpErrInfo.appendChild(tmpText);
                  let tmpLine = document.createElement("p");
                  tmpErrInfo.appendChild(tmpLine);
                  //tmpLine.style.borderBottom = "1px solid orange";
  
                  tmpLine.innerText = `Line: ${OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors[k].ErrorStrings[kk].Line}`;
                  tmpError.innerText = `Error: ${OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors[k].ErrorStrings[kk].Error}`;
                  tmpText.innerText = `Text: ${OUTPUT_RESULT.Projects[i].HTMLs[j].StrictVerify.Errors[k].ErrorStrings[kk].TextFromHTML}`;
                }
              }
            }
          } else {
            // no errors
            console.log("no errors")
            let errMsg = document.createElement("p");
            xhtmlErrorList.appendChild(errMsg);
            errMsg.innerText = "No errors!";
          }
        }
      } else {
        // no html. Show no html msg on screen
      }
    }
  }

}());
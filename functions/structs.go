package functions

import (
	"sync"
	"time"
)

// Work holds the data of everything
type Work struct {
	Mutex    sync.Mutex // For potential races when writing to Projects
	Projects []Project
	Complete bool
	Timing   Timing
}

// Project Holds one projects and its html file paths
type Project struct {
	Done       bool         // If project is verified
	FolderName string       // Name of folder
	HTMLs      []HTMLVerify // Holds that projects html files
	CSSs       []CSSVerify  // Holds css file(s) in project
}

// HTMLVerify holds one html doc and its warnings and errors
type HTMLVerify struct {
	Path         string       // Path to html doc
	AllVerified  bool         // If file has been verified for both strict and html5
	StrictVerify StrictVerify // For XHTML 1.0 Strict
	HTML5Verify  HTML5Verify  // For HTML5
}

// StrictVerify holds data for the XHTML 1.0 Strict verify
type StrictVerify struct {
	Verified bool // If file has been verified
	Result   string
	Errors   []ErrorGroup
	Warnings []string // All warnings
	Infos    []string
}

// HTML5Verify holds if file has warnings/errors for the HTML5. Should have none if ok
type HTML5Verify struct {
	Verified            bool // If file has been verified
	HasWarningsOrErrors bool // File has warnings
	ErrorValidating     error
}

// CSSVerify holds if css file(s) has warnings/errors for the CSS. Should have none if ok
type CSSVerify struct {
	Path        string // Path to css file
	Verified    bool   // If file has been verified
	HasWarnings bool   // File has warnings
	HasErrors   bool   // File has errors
}

// ErrorGroup holds an error group. Its type as string and the errors as []string
type ErrorGroup struct {
	ErrorType    string
	ErrorStrings []TheError
}

// TheError keeps the error data
type TheError struct {
	Line         string // The row and the title of the error
	Error        string
	TextFromHTML string // From the html doc itself
}

// Timing holds when the validation was started and when it was finished
type Timing struct {
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

package functions

import "sync"

// Work holds the data of everything
type Work struct {
	Mutex    sync.Mutex // For potential races when writing to Projects
	Projects []Project
	Complete bool
}

// Project Holds one projects and its html file paths
type Project struct {
	Done       bool         // If project is verified
	FolderName string       // Name of folder
	HTMLs      []HTMLVerify // Holds that projects html files
}

// HTMLVerify holds one html doc and its warnings and errors
type HTMLVerify struct {
	Path          string   // Path to html doc
	Verified      bool     // If file has been verified
	Warnings      []string // All warnings
	TotalWarnings int
	Errors        []string // All errors
	TotalErrors   int
}

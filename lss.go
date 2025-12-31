package main

import (
	"time"

	"github.com/fatih/color"
)

// Windows operating systems
const Windows = "windows"

// file types
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileText
	fileLink
)

// file extension
const (
	exe    = ".exe"
	deb    = ".deb"
	zip    = ".zip"
	tar    = ".tar"
	tar_gz = ".tar.gz"
	rar    = ".rar"
	png    = ".png"
	jpg    = ".jpg"
	jpeg   = ".jpeg"
	gif    = ".gif"
	txt    = ".txt"
	md     = ".md"
	csv    = ".csv"
	json   = ".json"
)

type file struct {
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	groupName        string
	size             int64
	modificationTime time.Time
	mode             string
}

type styleFileType struct {
	icon   string
	color  color.Attribute
	symbol string
}

var mapStyleByFileType = map[int]styleFileType{
	fileRegular:    {"📄", color.FgWhite, ""},
	fileDirectory:  {"📁", color.FgBlue, "/"},
	fileExecutable: {"⚙️", color.FgGreen, "*"},
	fileCompress:   {"🗜️", color.FgYellow, ""},
	fileImage:      {"🖼️", color.FgMagenta, ""},
	fileText:       {"📃", color.FgCyan, ""},
	fileLink:       {"🔗", color.FgBlue, "@"},
}

var (
	blue    = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	green   = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	red     = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	magenta = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()
	cyan    = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
)

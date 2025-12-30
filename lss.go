package main

import "time"

// operating systems
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
	color  string
	symbol string
}

var mapStyleByFileType = map[int]styleFileType{
	fileRegular:    {"📄", "white", ""},
	fileDirectory:  {"📁", "blue", "/"},
	fileExecutable: {"⚙️", "green", "*"},
	fileCompress:   {"🗜️", "yellow", ""},
	fileImage:      {"🖼️", "magenta", ""},
	fileText:       {"📃", "cyan", ""},
	fileLink:       {"🔗", "blue", "@"},
}

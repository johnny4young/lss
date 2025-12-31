package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/AJRDRGZ/fileinfo"
	"github.com/fatih/color"
	"golang.org/x/exp/constraints"
)

func main() {
	//filter pattern
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "show all files including hidden files")
	flagNumberRecords := flag.Int("n", 0, "number of records to generate")

	// order flags
	hasOrderByTime := flag.Bool("t", false, "order by time oldest to first")
	hasOrderBySize := flag.Bool("s", false, "sort by size smallest to largest")
	hasOrderReverse := flag.Bool("r", false, "reverse order")

	flag.Parse()

	path := flag.Arg(0)

	if path == "" {
		path = "."
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fs := []file{}

	for _, dir := range dirs {

		isHidden := isHidden(dir.Name(), path)

		if isHidden && !*flagAll {
			continue
		}

		if *flagPattern != "" {
			isMatched, err := regexp.MatchString("(?i)"+*flagPattern, dir.Name()) // case insensitive
			if err != nil {
				panic(err)
			}

			if !isMatched {
				continue
			}
		}

		f, err := getFile(dir, isHidden)
		if err != nil {
			panic(err)
		}

		fs = append(fs, f)
	}

	// ordering
	if !*hasOrderByTime || !*hasOrderBySize {
		orderByName(fs, *hasOrderReverse)
	}

	if *hasOrderBySize && !*hasOrderByTime {
		orderBySize(fs, *hasOrderReverse)
	}

	if *hasOrderByTime {
		orderByTime(fs, *hasOrderReverse)
	}

	if *flagNumberRecords == 0 || *flagNumberRecords > len(fs) {
		*flagNumberRecords = len(fs)
	}

	printList(fs, *flagNumberRecords)

}

func orderByName(file []file, isReverse bool) {
	sort.SliceStable(file, func(i, j int) bool {
		return mySort(strings.ToLower(file[i].name), strings.ToLower(file[j].name), isReverse)
	})
}

func orderBySize(file []file, isReverse bool) {
	sort.SliceStable(file, func(i, j int) bool {
		return mySort(file[i].size, file[j].size, isReverse)
	})
}

func orderByTime(file []file, isReverse bool) {
	sort.SliceStable(file, func(i, j int) bool {
		return mySort(file[i].modificationTime.Unix(), file[j].modificationTime.Unix(), isReverse)
	})
}

// generic implementation of sort
func mySort[T constraints.Ordered](i, j T, isReverse bool) bool {
	if isReverse {
		return i > j
	}
	return i < j
}

func printList(fs []file, nRecords int) {
	for _, file := range fs[:nRecords] {
		style := mapStyleByFileType[file.fileType]

		fmt.Printf("%s %s %s %10d %s %s %s%s %s\n", file.mode, file.userName, file.groupName,
			file.size, file.modificationTime.Format(time.DateTime), style.icon, setColor(file.name, style.color), style.symbol, markHidden(file.isHidden))
	}
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %v", err)
	}

	userName, groupName := fileinfo.GetUserAndGroup(info.Sys())

	f := file{
		name:             dir.Name(),
		fileType:         0,
		isDir:            dir.IsDir(),
		isHidden:         isHidden,
		userName:         userName,
		groupName:        groupName,
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(),
	}
	setFile(&f)

	return f, nil
}

func setFile(f *file) {
	switch {
	case isLink(*f):
		f.fileType = fileLink
	case f.isDir:
		f.fileType = fileDirectory
	case isExec(*f):
		f.fileType = fileExecutable
	case isCompress(*f):
		f.fileType = fileCompress
	case isImage(*f):
		f.fileType = fileImage
	default:
		f.fileType = fileRegular
	}
}

func setColor(nameFile string, styleColor color.Attribute) string {
	switch styleColor {
	case color.FgBlack:
		return blue(nameFile)
	case color.FgGreen:
		return green(nameFile)
	case color.FgRed:
		return red(nameFile)
	case color.FgMagenta:
		return magenta(nameFile)
	case color.FgYellow:
		return yellow(nameFile)
	case color.FgCyan:
		return cyan(nameFile)
	default:
		return nameFile
	}
}

func isLink(f file) bool {
	return strings.HasPrefix(strings.ToUpper(f.mode), "L")
}

func isExec(f file) bool {
	if runtime.GOOS == Windows {
		return strings.HasSuffix(strings.ToLower(f.name), exe)
	}
	return strings.Contains(f.mode, "x")
}

func isCompress(f file) bool {
	return strings.HasSuffix(f.name, tar_gz) ||
		strings.HasSuffix(f.name, tar) ||
		strings.HasSuffix(f.name, zip) ||
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, deb)
}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, jpeg) ||
		strings.HasSuffix(f.name, gif)
}

func isHidden(fileName, basePath string) bool {
	filepath := fileName
	if runtime.GOOS == Windows {
		filepath = path.Join(basePath, filepath)
	}

	return fileinfo.IsHidden(filepath)

}

func markHidden(isHidden bool) string {
	if !isHidden {
		return ""
	}

	return yellow("∅")
}

package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	//filter pattern
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "show all files including hidden files")
	flagNumberRecords := flag.Int("n", 0, "number of records to generate")

	// order flags
	//hasOrderByTime := flag.Bool("t", false, "order by time oldest to first")
	//hasOrderBySize := flag.Bool("s", false, "sort by size smallest to largest")
	//hasOrderReverse := flag.Bool("r", false, "reverse order")

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

		f, err := getFile(dir, false)
		if err != nil {
			panic(err)
		}
		fs = append(fs, f)
	}

	fmt.Println(fs)

	fmt.Println("pattern:", *flagPattern)
	fmt.Println("all:", *flagAll)
	fmt.Println("number of records:", *flagNumberRecords)
	//fmt.Println("order by time:", *hasOrderByTime)
	//fmt.Println("sort by size:", *hasOrderBySize)
	//fmt.Println("reverse order:", *hasOrderReverse)

}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %v", err)
	}
	f := file{
		name:             dir.Name(),
		fileType:         0,
		isDir:            dir.IsDir(),
		isHidden:         isHidden,
		userName:         "",
		groupName:        "",
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(),
	}

	return f, nil
}

package main

import (
	"flag"
	"fmt"
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

	fmt.Println("pattern:", *flagPattern)
	fmt.Println("all:", *flagAll)
	fmt.Println("number of records:", *flagNumberRecords)
	fmt.Println("order by time:", *hasOrderByTime)
	fmt.Println("sort by size:", *hasOrderBySize)
	fmt.Println("reverse order:", *hasOrderReverse)

}

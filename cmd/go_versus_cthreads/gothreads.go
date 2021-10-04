package main

import (
	"fmt"
	"os"
	"strconv"
)

func runFunction() {
	return
}

func main() {

	threadcount := 1

	if len(os.Args) > 1 {
		var ok error
		threadcount, ok = strconv.Atoi(os.Args[1])
		if ok != nil {
			fmt.Fprintf(os.Stderr, "Bad threadcount argument %q\n", threadcount)
			os.Exit(1)
		}
	}
	fmt.Printf("gothreads: threadcount=%v\n", threadcount)


	for i := 0; i < threadcount; i++ {
		go runFunction()
	}
}

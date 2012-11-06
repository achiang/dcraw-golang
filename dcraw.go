package main

import (
	"fmt"
	"flag"
)

var identify_only bool
func init() {
	flag.BoolVar(&identify_only, "i", false, "Identify files without decoding them")
}

func main() {
	flag.Parse()

	if len(flag.Arg(0)) == 0 {
		fmt.Println("No files to process.")
	}
	fmt.Println(flag.Arg(0))
}

package main

import (
	"fmt"
	"flag"
	"log"
	"os"
)

var identify_only bool
func init() {
	flag.BoolVar(&identify_only, "i", false, "Identify files without decoding them")
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No files to process.")
		os.Exit(1)
	}
	for _, arg := range flag.Args() {
		ifp, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 1)
		ifp.Read(buf)
	}
}

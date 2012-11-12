package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"encoding/binary"
	"bytes"
)

var identify_only bool
func init() {
	flag.BoolVar(&identify_only, "i", false, "Identify files without decoding them")
}

var order int16
var is_raw uint32
var ifp *os.File

func sget2(s []byte) int16 {
	var ret int16
	buf := bytes.NewBuffer(s)

	/* "II" means little-endian */
	if order == 0x4949 {
		binary.Read(buf, binary.LittleEndian, &ret)
	/* "MM" means big-endian */
	} else {
		binary.Read(buf, binary.BigEndian, &ret)
	}
	return ret
}

func get2() int16 {
	str := make([]byte, 2)
	ifp.Read(str)
	return sget2(str)
}

func sget4(s []byte) uint32 {
	var ret uint32
	buf := bytes.NewBuffer(s)
	if order == 0x4949 {
		binary.Read(buf, binary.LittleEndian, &ret)
	} else {
		binary.Read(buf, binary.BigEndian, &ret)
	}
	return ret
}

func get4() uint32 {
	str := make([]byte, 4)
	ifp.Read(str)
	return sget4(str)
}

func identify() {
	var hlen uint32
	order = get2()
	hlen = get4()
	fmt.Printf("hlen: %x\n", hlen)
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No files to process.")
		os.Exit(1)
	}
	for _, arg := range flag.Args() {
		var err error
		ifp, err = os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}
		identify()
	}
}

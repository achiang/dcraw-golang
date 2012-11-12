package main

import (
	"os"
	"testing"
)

func TestEndianness(t *testing.T) {
	var hlen uint32
	ifp, _ = os.Open("../../test/IMG_1100.CR2")

	if order != 0 {
		t.Errorf("order: %x, want: %x\n", order, 0)
	}
	if order = get2(); order != 0x4949 {
		t.Errorf("order: %x, want: %x\n", order, 0x4949)
	}
	/* Value obtained by inspection of original dcraw.c using gdb */
	if hlen = get4(); hlen != 0x10002a {
		t.Errorf("hlen: %x, want: %x\n", hlen, 0x10002a)
	}
	ifp.Close()
}

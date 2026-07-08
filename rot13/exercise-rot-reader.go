package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case b >= 65 && b <= 77 || b >= 97 && b <= 109:
		return b + 13
	case b >= 78 && b <= 90 || b >= 110 && b <= 122:
		return b - 13
	default:
		return b
	}
	return 0
}

func (r13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

/*
Input 	ABCDEFGHIJKLM NOPQRSTUVWXYZ abcdefghijklm nopqrstuvwxyz
Output 	NOPQRSTUVWXYZ ABCDEFGHIJKLM nopqrstuvwxyz abcdefghijklm
*/
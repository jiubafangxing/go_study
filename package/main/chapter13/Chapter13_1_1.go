package chapter13

import (
	"io"
	"os"
)

func Chapter1311Test() {
	f, _ := os.Open("./READ_FILE.log")
	defer f.Close()
	var r io.Reader = f
	for {
		buf := make([]byte, 512)
		_, err := r.Read(buf)
		if nil != err {
			break
		}

	}
}

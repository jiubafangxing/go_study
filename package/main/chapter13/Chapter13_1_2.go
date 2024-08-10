package chapter13

import (
	"bufio"
	"io"
	"os"
)

func Chapter1312Test() {
	f, _ := os.Open("./READ_FILE.log")
	defer f.Close()
	var r io.Reader = f
	var br *bufio.Reader = bufio.NewReaderSize(r, 1024)
	for {
		buf := make([]byte, 512)
		_, err := br.Read(buf)
		if nil != err {
			break
		}

	}
}

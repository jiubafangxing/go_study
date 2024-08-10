package chapter13

import (
	"bytes"
	"io"
	"log"
)

func Chapter1317Test() {
	data := []byte{1, 2, 3, 4, 5}
	r := bytes.NewReader(data)
	sr := io.NewSectionReader(r, 2, 2)
	for {
		st := make([]byte, 1)
		n, err := sr.Read(st)
		if nil != err {
			break
		}
		log.Println(n, err, st)

	}
}

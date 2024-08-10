package chapter13

import (
	"bytes"
	"io"
	"log"
)

func Chapter1316Test() {
	data := []byte{1, 2, 3, 4, 5}
	r := bytes.NewReader(data)
	l := io.LimitedReader{r, 2}
	for {
		bs := make([]byte, 1)
		n, err := l.Read(bs)
		if nil != err {
			break
		} else {
			log.Println(bs, n, err)
		}
	}
}

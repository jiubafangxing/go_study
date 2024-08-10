package chapter13

import (
	"bytes"
	"log"
)

func Chapter1313Test() {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte{1, 2, 3})
	log.Println(buf.Bytes())
}

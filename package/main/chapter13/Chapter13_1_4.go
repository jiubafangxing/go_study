package chapter13

import (
	"bytes"
	"io"
	"log"
)

func Chapter1314Test() {
	data := []byte{1, 2, 3, 4, 5, 6}
	r := bytes.NewReader(data)
	for i := 0; i < 3; i++ {
		if i == 2 {
			r.Seek(-2, io.SeekEnd)
		}
		dest := make([]byte, 1)
		n, err := r.Read(dest)
		if n > 0 && nil == err {
			log.Println(dest)
		}
	}
	buf := make([]byte, 1)
	r.ReadAt(buf, 1)
	log.Println(buf)
	b, err := r.ReadByte()
	if nil == err {
		log.Println(b)
	}
}

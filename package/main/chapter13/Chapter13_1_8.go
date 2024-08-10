package chapter13

import (
	"bytes"
	"io"
	"log"
)

func Chapter1318Test() {
	data1 := []byte{1, 2, 3, 3, 4, 5, 6}
	data2 := []byte{6, 7, 8, 4, 2, 2}
	d1r := bytes.NewReader(data1)
	d2r := bytes.NewReader(data2)
	w1 := bytes.NewBuffer(nil)
	w2 := bytes.NewBuffer(nil)
	mr := io.MultiReader(d1r, d2r)
	mw := io.MultiWriter(w1, w2)
	for {
		b2 := make([]byte, 2)
		n, err := mr.Read(b2)
		log.Println(b2)
		if nil != err {
			break
		} else {
			log.Println(n, err)
		}
		mw.Write(b2[:n])
	}
	log.Println(w1.Bytes())
	log.Println(w2.Bytes())
}

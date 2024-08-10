package chapter13

import (
	"bytes"
	"log"
)

func Chapter1315Test() {
	data := []byte{1, 2, 3, 2, 3, 4, 2}
	r := bytes.NewReader(data)
	n, err := r.ReadByte()
	if nil != err {
		log.Fatalln(err)
	} else {
		log.Println(n)
	}
	n, err = r.ReadByte()
	if nil != err {
		log.Fatalln(err)
	} else {
		log.Println(n)
	}
	err = r.UnreadByte()
	err = r.UnreadByte()
	if nil != err {
		log.Fatalln(err)
	}
	n, err = r.ReadByte()
	if nil != err {
		log.Fatalln(err)
	} else {
		log.Println(n)
	}
}

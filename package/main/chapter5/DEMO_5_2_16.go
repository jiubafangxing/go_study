package main

import (
	"errors"
	"log"
)

func test() {
	stack := make([]int, 0, 5)
	push := func(a int) error {
		l := len(stack)
		if l == cap(stack) {
			return errors.New("out of capicity")
		}
		stack = stack[:l+1]
		stack[l] = a
		return nil
	}
	pop := func() (int, error) {
		l := len(stack)
		if l == 0 {
			return -1, errors.New("no elements ")
		}
		result := stack[l-1]
		stack = stack[:l-1]
		log.Println("cap size is ", cap(stack))
		return result, nil
	}
	err := push(1)
	if nil != err {
		log.Fatalln(err)
	}
	log.Println(len(stack))
	popResult, err2 := pop()
	if nil != err2 {
		log.Fatalln(err2)
	}
	log.Println(popResult)
	err = push(3)
	if nil != err {
		log.Fatalln(err)
	}

}

func main() {
	test()
}

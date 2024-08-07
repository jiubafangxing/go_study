package chapter8

import (
	"log"
	"sync"
)

type DEMO_8_1_7data [8]struct {
	id     int
	result int
}

func DEMO_8_1_7test() {
	var wg sync.WaitGroup
	var data1 DEMO_8_1_7data
	//simulate local storage
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(x int) {
			data1[x].id = x
			data1[x].result = x
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Println("%+v\n", data1)
}

func DEMO_8_1_7main() {
	DEMO_8_1_7test()
}

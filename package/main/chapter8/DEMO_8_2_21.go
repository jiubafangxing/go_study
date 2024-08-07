package chapter8

import "log"

const (
	max     = 50000000
	block   = 500
	bufsize = 100
)

func DEMO_8_2_21test() {
	done := make(chan struct{})
	ch1 := make(chan int, bufsize)
	go func() {
		defer close(done)
		count := 0
		for x := range ch1 {
			count += x
		}
		log.Printf("result is %d \n ", count)
	}()
	for i := 0; i < max; i++ {
		ch1 <- i
	}

	close(ch1)
	<-done
}

func DEMO_8_2_21test1() {
	done := make(chan struct{})
	ch1 := make(chan [block]int, bufsize)
	j := 0
	go func() {
		defer close(done)
		count := 0
		for block := range ch1 {
			for _, x := range block {
				count += x
			}
		}
		log.Printf("result is %d \n ", count)
	}()
	batch := [block]int{}
	for i := 0; i < max; i++ {
		if j < block {
			batch[j] = i
			j++
		}
		if j == block || i == max-1 {
			ch1 <- batch
			j = 0
			batch = [block]int{}
		}
	}

	close(ch1)
	<-done
}

func DEMO_8_2_21main() {
	DEMO_8_2_21test()
}

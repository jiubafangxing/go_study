package chapter1

import "time"

func taska(id int) {
	for i := 0; i < 5; i++ {
		println("id", id, "value", i)
		time.Sleep(time.Second)
	}
}

func goroutine_demomain() {
	go taska(1)
	go taska(100)
	time.Sleep(time.Second * 6)
}

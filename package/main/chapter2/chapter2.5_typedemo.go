package chapter2

type myt byte

const (
	chapter2a myt = 1 << iota
	chapter2b
	chapter2c
)

func chapter25_typedemomain() {
	println(chapter2a)
	println(chapter2b)
	println(chapter2c)
}

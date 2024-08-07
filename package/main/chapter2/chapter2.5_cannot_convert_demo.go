package chapter2

func chapter25_cannot_convert_demomain() {
	a := 1
	p := (*int)(&a)
	println(*p)
}

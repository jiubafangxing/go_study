package chapter1

func div1(a, b int) int {
	defer println("call dev", a, b)
	return a / b
}

func defer_demomain() {
	div1(1, 2)
	div1(1, 0)
}

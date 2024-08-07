package chapter1

import "strconv"

func chatper24_strconv_demomain() {
	a, _ := strconv.ParseInt("010010", 2, 10)
	println(a)
}

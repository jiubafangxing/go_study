package chapter1

func map_demomain() {
	mymap := make(map[string]int)
	mymap["a"] = 1
	k, exist := mymap["b"]
	println(k, exist)
	k1, v1 := mymap["a"]
	println(k1, v1)
}

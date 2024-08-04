	
package main
func test(){
	m :=map[string]int{
		"a":1,
		"b":2,
	}
	m["a"] = 4
	m["c"] =1
	if v,ok :=m["d"]; ok{
		println(v)
	}
	delete(m, "d")
}

func main(){
	test()
}

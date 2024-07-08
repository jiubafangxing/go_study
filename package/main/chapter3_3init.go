package main

func main(){
	type data  struct{
		name string
		age int
	}
	i := data{
		"kao",
		1,
	}
	println(i.name)
}



package chapter3

func chapter3_3initmain(){
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



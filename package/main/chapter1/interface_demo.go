package main

type Animal interface {
	Eat()
}

type Cat struct {
}

func (Cat) Eat() {
	println("cat eat fish")
}

func main() {
	var cat Cat
	var an Animal = cat
	an.Eat()
}

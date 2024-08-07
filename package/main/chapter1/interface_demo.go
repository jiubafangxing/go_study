package chapter1

type Animal interface {
	Eat()
}

type Cat struct {
}

func (Cat) Eat() {
	println("cat eat fish")
}

func interface_demomain() {
	var cat Cat
	var an Animal = cat
	an.Eat()
}

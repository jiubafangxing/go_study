package chapter3

func count() int {
	println("count")
	return 3
}
func chapter33_funcformain() {
	for x := count(); x < 10; x++ {
		println(x)
	}
	for x := 1; x < count(); x++ {
		println(x)
	}
}

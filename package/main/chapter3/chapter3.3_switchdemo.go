package chapter3

func chapter33_switchdemomain() {
	a := 2
	switch a {
	case 1:
		print(1)
	case 2:
		println(2)
		fallthrough
	case 3:
		println("fallthrough")
	}
}

package chapter3

func chapter33_switch_replace_ifmain() {
	switch i := 1; {
	case i == 1:
		println(1)
	case i != 1:
		println("not 1")
	}
}

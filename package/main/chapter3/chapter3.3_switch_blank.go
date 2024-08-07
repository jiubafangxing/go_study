package chapter3

func chapter33_switch_blankmain() {
	a, b := 2, 3
	switch i := 1; i {
	case a:
	case b:
	default:
		println(i)
	}
}

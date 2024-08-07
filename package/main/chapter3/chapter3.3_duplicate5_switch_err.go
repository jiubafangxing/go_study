package chapter3

func chapter33_duplicate5_switch_errmain() {
	switch i := 1; i {
	//case 5:
	case 5, 6:
		println(i)
	}
}

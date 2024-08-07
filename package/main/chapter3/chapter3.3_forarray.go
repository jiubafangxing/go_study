package chapter3

func chapter33_forarraymain() {
	strs := [3]string{"1", "a", "b"}
	for i, v := range strs {
		println(i, v)

	}
}

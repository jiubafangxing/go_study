package chapter2

func chapter2_6short_var3main() {
	x := 1
	x, y := 2, 2
	println(x, y)
	{
		x, y := 3, 3
		println(x, y)
	}
	println(x, y)
}

package chapter1

func slice_demomain() {
	array := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		array = append(array, i)
	}
	for i := 0; i < len(array); i++ {
		println(array[i])
	}
}

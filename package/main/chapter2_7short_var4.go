package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/degv/a")
	fmt.Println(f, err)
	buf := make([]byte, 1024)
	d, err := f.Read(buf)
	fmt.Println(d, err)
}

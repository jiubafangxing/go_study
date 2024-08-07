package chapter2

import (
	"fmt"
	"os"
)

func chapter2_7short_var4main() {
	f, err := os.Open("/degv/a")
	fmt.Println(f, err)
	buf := make([]byte, 1024)
	d, err := f.Read(buf)
	fmt.Println(d, err)
}

package main
import (
"fmt"
"math"
)

func main(){
	a, b, c := 0x11, 014,100
	fmt.Println(a,b,c)
	fmt.Printf("0b%b, %#o, %#x", c,c,c)
	fmt.Println(math.MinInt8)
}

package main
import(
	"errors"
)
func div(a, b int)(int ,error){
	if b< 0{
		return 0, errors.New("can not be zero")
	}
	return a/b, nil
}

func main(){
	x,y := div(1,2)
	println(x,y)
	c,d := div(1,0)
	println(c,d)
}

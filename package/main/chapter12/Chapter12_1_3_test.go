package chapter12

import(
	"os"
	"testing"
)

func Add1213(a,b int)(int){
	return a+b
}

func Minus1213(a,b int)(int){
	return a-b
}

func TestAdd1213(t *testing.T){
	t.Parallel()
	if 3 != Add1213(1,2) {
		t.FailNow()
	}
}

func TestMinus1213(t *testing.T){
	if os.Args[len(os.Args)-1] == "b"{
		t.Parallel()
	}
	if 3 != Minus1213(4,1){
		t.FailNow()
	}

}

package chapter12
import ("testing")
func Add1211(a,b int)(int){
	return  a+b
}
func TestAdd1211(t *testing.T){
	if (3 != Add1211(1,2)){
		t.FailNow()
	}
} 

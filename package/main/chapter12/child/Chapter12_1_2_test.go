package child

import("testing")

func Add1212(a,b int)(int){
	return a + b
}

func TestAdd1212(t *testing.T){
	if(Add1212(1,2) != 3){
		t.FailNow()
	}
}

package main

func div(a , b int)(int ){
  	defer println("call dev", a,b)	
	return a/b;
}

func main(){
	div(1,2)
	div(1,0)
}

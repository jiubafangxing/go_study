package main

import (
	"errors"
	"fmt"
)

func div(x,y int) (int , error){
	if(y != 0){
		return x/y, nil
	}
	return 0, errors.New(" division by zero")
}

func main(){
	result, err := div(4,2)
	if(nil == err){
		print("the result is ", result);
	}
	result, err  = div(4,0)
	if(nil != err){
		fmt.Println(err)
	}
}

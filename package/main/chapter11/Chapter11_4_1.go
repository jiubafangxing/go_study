package chapter11
import (
	"log"
	"reflect"
	"strings"
)

func add(args []reflect.Value)(results []reflect.Value){
	if len(args) == 0{
		return nil
	}
	var ret  reflect.Value
	switch args[0].Kind(){
	case reflect.Int:
		n := 0 
		for _, v := range args{
			n+=int(v.Int())	
		}
		ret = reflect.ValueOf(n)

	case reflect.String:
		ss := make ([]string, 0, len(args))
		for _,v := range args{
			ss = append(ss, v.String())
		}
		ret = reflect.ValueOf(strings.Join(ss, ""))

	}
	results =  append(results, ret)
	return results
}

func Chapter1141Test(){
	var addInt func(x,y int)int
	var addStr func(x1,y1 string)string
	makeAdd(&addInt)
	makeAdd(&addStr)
	log.Println(addInt(100,200))
	log.Println(addStr("100","200"))
}


func makeAdd(fptr interface{}){
	aiv := reflect.ValueOf(fptr)
	aie := aiv.Elem()
	r := reflect.MakeFunc(aie.Type(),add)
	aie.Set(r)
}

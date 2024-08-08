package chapter11
import (
"fmt"
"reflect"
"log"
)
type user struct {
	age int
	name string
}
type manager struct{
	level int
	user
}

func (m manager) Stringer()(string){
	return fmt.Sprintf("age:%d, name:%s, level:%d",m.age, m.name,m.level )
}

func Chapter1115Test(){
	m := manager{
		level:1,
	}
	m.user.name = "john"
	m.user.age = 12
	mp := &m
	mptype :=reflect.TypeOf(mp)	
	if mptype.Kind() == reflect.Ptr{
		mptype = mptype.Elem()
	}
	log.Println(mptype)
	nf := mptype.NumField()
	for i:=0;i< nf;i++{
		f := mptype.Field(i)
		if f.Anonymous{
			anf := f.Type.NumField()
			if anf > 0{
				for ai:=0;ai< anf;ai++{
					af := f.Type.Field(ai)
					log.Printf("f name :%s, f type :%s", af.Name,af.Type )
				}
			}
		}else{

					log.Printf("f name :%s, f type :%s", f.Name,f.Type )

		}
	}

}

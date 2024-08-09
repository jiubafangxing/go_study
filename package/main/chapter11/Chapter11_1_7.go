package chapter11

import (
	"log"
	"reflect"
)

func (m *manager) PrintManager() {
	log.Println(m.level)
}
func Chapter1117Test() {
	m := manager{}
	mt := reflect.TypeOf(&m)
	s := []reflect.Type{mt, mt.Elem()}
	for _, t := range s {
		nm := t.NumMethod()
		for j := 0; j < nm; j++ {
			tm := t.Method(j)
			log.Printf("type : %s, method name : %s", t.Name(), tm.Name)
		}
	}
}

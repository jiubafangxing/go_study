	
package chapter8
import(
	"log"
)
type pool chan  []byte

func  NewPool(cap int)pool{
	return pool(make(chan []byte,cap))
}

func (p pool) get()(v []byte){
	select {
	case v	=<- p:
	default :
		v = make([]byte, 10)
	}
	return v
}

func  (p pool) put(b []byte){
	select{
		case p <- b:
		default:
	}

}
func DEMO_8_2_17test(){
	var p1 pool = NewPool(3)
	log.Println(p1)
}

func DEMO_8_2_17main(){
	DEMO_8_2_17test()
}

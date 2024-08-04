	
package main
import(
	"log"
)
func test(){
	d :=[...]int{0,10,10,20,120,301,320,120,123,120,12}
	s1 := d[3:7]
	s2 :=s1[1:6]
	s2[3] =333
	log.Printf("%#v", d)
	log.Printf("%#v", s1)
	log.Printf("%#v", s2)

}

func main(){
	test()
}

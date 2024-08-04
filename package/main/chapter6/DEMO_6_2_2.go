	
package main
import(
	"log"
)
type user struct{}

type manager struct{
	user
}

func (u user) toString()string{
	return "u"
}

func (m manager) toString()string{
	return m.user.toString()+"m"
}

func test(){
	m := manager{}
	log.Println(m.toString())
}

func main(){
	test()
}

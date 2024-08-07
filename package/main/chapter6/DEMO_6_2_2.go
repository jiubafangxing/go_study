	
package chapter6
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

func DEMO_6_2_2test(){
	m := manager{}
	log.Println(m.toString())
}

func DEMO_6_2_2main(){
	DEMO_6_2_2test()
}

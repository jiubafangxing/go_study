	
package chapter6
import(
	"log"
)

type User struct{
	age int
}
func (u User)changeAge(){
	u.age +=1
}
func DEMO_6_4_1test(){
	u := User{
		age:1,
	}
	u.changeAge()
	log.Println(u.age)
	(&u).changeAge()
	log.Println(u.age)
	cCall := User.changeAge
	cCall(u)
	log.Println(u.age)


}

func DEMO_6_4_1main(){
	DEMO_6_4_1test()
}

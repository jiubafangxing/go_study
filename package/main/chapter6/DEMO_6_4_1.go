	
package main
import(
	"log"
)

type User struct{
	age int
}
func (u User)changeAge(){
	u.age +=1
}
func test(){
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

func main(){
	test()
}

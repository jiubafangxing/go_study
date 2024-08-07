package chapter9
import(
	"log"
)
func chapter916print [T int|string] (t T){
	switch any(t).(type){
		case int: log.Printf("数字%d\n", t)
		case string: log.Printf("string%s\n", t)
			
	}
}
func Chapter916Test() {
	chapter916print(1)
	chapter916print("hello")
	return
}

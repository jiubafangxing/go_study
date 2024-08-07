	
package main
import(
	"log"
)

const(
	max = 50000000
	block = 500
	bufsize = 100
)
func test(){
	done := make(chan struct{})
	ch1 := make(chan int, bufsize)
	for i:=0;i< max;i++{
		ch1 <- i
	}
	go func(){
		defer close(done)
		count :=0
		for x:= range <- ch1{
			count += x
		}
	}()
	close(ch1)
	<- done
}

func test1(){
	done := make(chan struct{})
	ch1 := make(chan [block]int, bufsize)
	for i:=0;i< max;i++{
		//TODO compare batch
	}
}

func main(){
	test()
}

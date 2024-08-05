	
package main

type N int

type Stringer interface{
	toString()(string)
}
func (N) toString()string{
	return "a"
}

func init(){
	var _ Stringer = N(0)
}
func test(){

}

func main(){
	test()
}

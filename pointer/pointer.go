package main
import "fmt"
func val1(ival int) {
	ival=0
}

func val2(iptr *int) {
	*iptr=0
}

func main() {
	i:=1
	fmt.Println("initial",i)
	val1(i)
	fmt.Println("val1",i)
	val2(&i)
	fmt.Println("val2",i)
	fmt.Println("pointer add",&i)
}
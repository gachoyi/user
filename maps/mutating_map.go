package main
import "fmt"
func main() {
	m:=make(map[string]int)
	m["answer"]=42
	fmt.Println("the value:",m["answer"])
	elem := m["answer"]
	fmt.Println(elem)
	m["answer"]=48
	fmt.Println("the value:",m["answer"])
	
	delete(m,"answer")
	fmt.Println("the value",m["answer"])

	//m["answer"]=12
	a, ok :=m["answer"]
	fmt.Println("the value:",a,"present?",ok)
}		
package main
import "fmt"
type person struct {
	name string
	age int
}


func main() {
	fmt.Println(person{"bob",20})
	s:=person{"mike",21}
	fmt.Println(s.name)
	sp:=&s
	sp.age=50
	fmt.Println(s.age)
}
package main
import "fmt"
type rec struct {
	width,height int
}

func (r *rec) area() int {
	return r.width*r.height
}

func (r rec) perim() int {
	return 2*r.width+2*r.height
}

func main() {
	r:=rec{width:10,height:5}
	fmt.Println("area",r.area())
	fmt.Println("perim",r.perim())
	rp:=&r
	fmt.Println("area",rp.area())
	fmt.Println("perim",rp.perim())
}
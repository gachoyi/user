package main
import (
	"fmt"
	
)

func main() {
	fmt.Println(len("hello,世界"))
	fmt.Println(len([]rune("hello,世界")))
}

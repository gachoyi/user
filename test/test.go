package main
import (
	"fmt"
	
)

func main() {
	fmt.Println(len("hello,world"))
	fmt.Println(len([]rune("hello,world")))
}

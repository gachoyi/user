package main
import (
	"fmt"
	"math"
)

func main() {
	const s string = "constant"
	fmt.Println(s)
	const n = 5000000
	const d = 3e20/n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
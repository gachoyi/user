package main
import (
    "fmt"
)
var (
    i uint64 = 1<<4-1
)
func main() {
    d := []byte{'r', 'o', 'a', 'd'}
    e := d[2:]
    e[1] = 'm'
    g := 2<<3
    const f = "%T(%v)\n"
    fmt.Printf(f,d,d)
    fmt.Printf(f,e,e)
    fmt.Println(g,i)
}

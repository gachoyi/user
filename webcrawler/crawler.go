package main
import (
	
	"fmt"
	"flag"
	"os"
	"net/http"
	"io/ioutil"
)


func main() {
	flag.Parse()       //parse命令后面的参数到slice。
	args := flag.Args()    //args可能是空的，也可能有若干个string值
	
	if len(args) < 1 {     //判断一下args是不是空的，如果是空就退出。
		fmt.Println("please specify start page")
		os.Exit(1)
	}
}
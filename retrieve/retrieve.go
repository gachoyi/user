package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main()  {
	resp, err:=http.Get("http://www.cssf.cc")
	fmt.Println("http error is:",err)
	body,err := ioutil.ReadAll(resp.Body)
	fmt.Println("read error is:",err)
	fmt.Println(string(body))
}
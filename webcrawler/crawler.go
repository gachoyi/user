package main
import (
	"crypto/tls"
	"fmt"
	"flag"
	"os"
	"net/http"
	"net/url"
	"github.com/JackDanger/collectlinks"
)

var visited = make(map[string]bool)     //visited这个dictionary用于存储url->bool对。 bool用于表示是否已经访问过。


func main() {
	flag.Parse()                                                   //parse命令后面的参数到slice。
	args := flag.Args()                                       //args可能是空的，也可能有若干个string值
	
	if len(args) < 1 {                                       //判断一下args是不是空的，如果是空就退出。
		fmt.Println("please specify start page")
		os.Exit(1)
	}
	
	queue:=make(chan string)       //创建一个chan
	go func() {
		queue<- args[0]           //命令行的第一个参数放到chan queue
	}()
	for uri :=range queue{
		enqueue(uri,queue)          //  对chan中的每项执行func enqueue
	}
	
}



func enqueue(uri string, queue chan string) {
	fmt.Println("fetching",uri)
	visited[uri] = true      //将要访问的uri放进visited，并赋值为true
	tlsConfig := &tls.Config{                             //&的意思：生成一个tls.Config对象，其InsecureSkipVerify值设为true
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig : tlsConfig,
	}
	client := http.Client{Transport: transport}
	resp,err := client.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()              //defer等上面的client.Get执行完毕后关闭连接
	links :=collectlinks.All(resp.Body)           //links是一个包含了二级超链接的slice
	for _,link := range links {
		absolute:=fixUrl(link,uri)
		if uri != "" {
			if !visited[absolute] {  //如果absolute在visited里不为true，就将absolute放进queue channle
				go func() {queue<-absolute}()   
			}   //二级超链接下可能还有下级链接，所以放到queue中继续执行enqueue。
		}
		
	}
	
}

func fixUrl(href,base string) (string) {
	uri,err:= url.Parse(href)  //用url.Parse()将rawurl转换为URL structure
	//rawurl是domain后面的一部分URL，URL “http://www.xyz.com/articles/recent.aspx”的rawurl部分是“/articles/recent.aspx”
	if err!=nil{
		return ""
	}
	baseUrl,err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)  //ResolveReference根据一个baseUrl解析出相对url（uri）的绝对地址。
	                                    //如果uri已经是一个绝对地址就不做任何修改，直接返回这个地址
	return uri.String()
}














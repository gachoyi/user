package main
import "fmt"
func main() {
	s:=make([]string,3)
	fmt.Println("empty",s)
	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println("set",s)
	fmt.Println("get",s[2])
	fmt.Println("length",len(s))
	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("appended",s)
	
	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("copy of s",c)
	l:=s[2:5]
	fmt.Println("slice1",l)
	l=s[:2]
	fmt.Println("slice2",l)
	l=s[2:]
	fmt.Println("slice3",l)
	t:=[]string{"g","h","i"}
	fmt.Println("dc1",t)
	
	twoD := make([][]int,3)
	for i:=0;i<3;i++ {
		innerlen:= i+1
		twoD[i]=make([]int,innerlen)
		for j:=0;j<innerlen;j++ {
			twoD[i][j]=i+j
		}
	}
	fmt.Println("twoD",twoD)
}
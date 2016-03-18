package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		cs <- cakeName
	}
	<-cs
	//close(cs)
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
	strbry_closed, choco_closed := false, false
	for {
		if (strbry_closed && choco_closed) {return}
		fmt.Println("Waiting for a new cake...") 
		select {
		case cakeName, strbry_ok := <-strbry_cs:
			if  (!strbry_ok)  {
				strbry_closed = true
				fmt.Println("strawberry channel closed")
			} else {
				fmt.Println("received from strawberry channel, now packing", cakeName)
			}
		case cakeName, choco_ok := <-choco_cs:
			if  (!choco_ok)  {
				choco_closed = true
				fmt.Println("chocho channel closed")
			} else {
				fmt.Println("received from choco channel,now packing", cakeName)
			}
		}
	}
}

func main() {
	strbry_cs := make(chan string)
	choco_cs := make(chan string)
	go makeCakeAndSend(strbry_cs, "strawberry", 3)
	go makeCakeAndSend(choco_cs, "choco", 3)
	
	go receiveCakeAndPack(strbry_cs, choco_cs)
	time.Sleep(2 * 1e9)
}

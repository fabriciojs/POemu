package main

import "fmt"
import "time"

func main() {
	ch := make(chan bool)
	ch2 := make(chan int)
	
	go func() {
		time.Sleep(time.Second * 5)
		ch <- true
	}()
	var a interface{}
	select {
	case a = <- ch : { }
	case a = <- ch2 : { }
	}
	
	fmt.Println("finished")
	fmt.Println(a)
}

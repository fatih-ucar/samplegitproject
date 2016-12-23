package main

import "fmt"

var delegate chan func(msg string)
var c chan bool

func worker(f func(msg string)) {
	fmt.Println("worker1")
	delegate <- f
	fmt.Println("worker2")

}

func runWorker(msg string) {
	fmt.Println("runworker1")
	f := <-delegate
	fmt.Println("runworker2")
	f(msg)
	c <- true

}

func init(){
	fmt.Println("git rename")
}

func main() {
	c = make(chan bool)
	delegate = make(chan func(msg string))
	go worker(func(msg string) {
		fmt.Printf("%v bu benim birinci str", msg)
	})
	go runWorker("fatih")

	<-c
}

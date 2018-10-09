package main

import "fmt"

func sum(s []int, c chan<- int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	var d <-chan int = c //channel is a pointer type that point to a channel data
	var e chan<- int = c

	go sum(s[:len(s)/2], e)
	go sum(s[len(s)/2:], e)
	x, y := <-d, <-d // 从 c 中接收

	fmt.Println(x, y, x+y)
}

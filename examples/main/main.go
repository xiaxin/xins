package main

import "fmt"

type T struct{}

func IsClosed(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func main() {
	c := make(chan struct{})
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
	fmt.Println(IsClosed(c)) // true
}

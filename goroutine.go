package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- "Goroutine : " + strconv.Itoa(i)
				
			}
		}(i)
	}

	for k := 1; k <= 10; k++ {
		fmt.Println(k, <-ch)
		time.Sleep(1*time.Second)
	}
}
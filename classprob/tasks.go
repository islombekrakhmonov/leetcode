package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var second = make(chan int)
var output []int
var sum int
// var wg = sync.WaitGroup{}
var urls = []string{
	"https://google.com",
    "https://tutorialedge.net",
    "https://twitter.com",
}
func main() {
	// wg.Add(1)
	// // go firstTast()
	// go secondTask()
	// wg.Wait()
	secondTask()

}


// func firstTast(){

// 	func (){
// 		for i:=1;i<=10;i+=2{
// 				fmt.Println("Odd Numbers:",i)
// 				time.Sleep(1*time.Second)
// 			}
// 	}()

// 	func (){
// 		for i:=1;i<=10;i++{
// 			if i%2==0{
// 				fmt.Println("Even Numbers:",i)
// 				time.Sleep(1*time.Second) 
// 			}
// 		}
// 	}()
// 	wg.Done()
// }
// 2.Write a program that creates a Go routine to read integers from standard input and send them to a channel. Create another Go routine that reads from the channel and prints the sum of the integers received.
func secondTask(){
	var inputNumber1 = func(){ 
		for {
			var number int 
			fmt.Println("Please input your number")
			_, err := fmt.Scan(&number)
			if err != nil{
				close(second)
				return
			}
			second <- number
		}
		} 
		channel := func ()  {
			for num := range second{
				sum +=num
			}
			fmt.Println(sum)
		}
		go inputNumber1()
		channel()
	}


















type result struct{
	url string
	err error
	latency time.Duration
}

func get(url string, ch chan<-result){
	start := time.Now()

	if resp, err := http.Get(url); err != nil{
		ch <- result{url, err, 0}
	} else {
		t:= time.Since(start).Round(time.Millisecond)
		ch<- result{url,nil,t}
		resp.Body.Close()
	}

}

func lil (){
	results := make(chan result)
	

	for _,url := range urls{
		go get(url,results)
	}

	for range urls{
		r:= <-results

		if r.err != nil{
		log.Printf("%-20s %s\n", r.url, r.err)
	} else {
		log.Printf("%-20s %s\n", r.url, r.latency)
	}

	}

}
	




// done := make(chan bool,3)
// for i:=0;i<3;i++{
// 	go worker(i,done)
// }

// for i:= 0;i<5;i++{
// 	tasks<- "Task"
// }
// close(tasks)

// for i:=0;i<3;i++{
// 	<-done
// }




// func worker(id int, done chan bool){
// 	for {
// 		task, more := <-tasks
// 		if more {
// 			fmt.Println("worker", id, "processing task", task)
// 		} else {
// 			done <- true
// 			return
// 		}
// 	}
// }



// func input(){
// 	var inputInfo string
// 	fmt.Scan(&inputInfo)
// 	go func(){pop <- inputInfo}()	
// }

// func receives(receiver chan<- string){

	
// }

// 1.Write a program that creates two Go routines, one to print even numbers from 1 to 10 and the other to print odd numbers from 1 to 10.

// 2.Write a program that creates a Go routine to read integers from standard input and send them to a channel. Create another Go routine that reads from the channel and prints the sum of the integers received.

// 3.Write a program that creates a Go routine to generate random numbers between 1 and 100 and send them to a channel. Create another Go routine that reads from the channel and prints the numbers in descending order.

// 4.Write a program that creates a Go routine to download files from a list of URLs. Use a channel to communicate the progress of the downloads. Print out the percentage completed for each download.

// 5.Write a program that creates a Go routine to calculate the factorial of a number using recursion. Use a channel to communicate the result of the calculation to the main program.
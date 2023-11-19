package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func sum(s []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum

	runtime.UnlockOSThread()
}

func main() {
	// s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int, 2) // Use a buffered channel to avoid deadlock
	// var wg sync.WaitGroup

	// wg.Add(2)

	// go func() {
	// 	sum(s[:len(s)/2], c, &wg)
	// }()

	// go func() {
	// 	sum(s[len(s)/2:], c, &wg)
	// }()

	// wg.Wait()
	// close(c)

	// x := <-c
	// y := <-c

	// fmt.Println(x, y, x+y)

	// msth := make(chan string, 2)

	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		msth <- fmt.Sprintf("number, %v", i)
	// 	}
	// 	close(msth)
	// }()

	// for {
	// 	v, ok := <-msth
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println("Received ", v, ok)
	// }

	// fmt.Println("Closed")

	// go func() {
	// 	for {
	// 		time.Sleep(time.Second * 1)
	// 		msth1 <- "second"
	// 	}
	// }()

	// for {
	// 	select {
	// 	case msg := <-msth:
	// 		fmt.Println(msg)
	// 	case msg := <-msth1:
	// 		fmt.Println(msg)
	// 		default :
	// 		fmt.Println("NOthing")
	// 	}
	// }

	// ch := make(chan int, 2)
	// go write(ch)
	// time.Sleep(2 * time.Second)
	// for v := range ch {
	// 	fmt.Println("read value", v, "from ch")
	// 	time.Sleep(2 * time.Second)

	// }

	// startTime := time.Now()
	// noOfJobs := 100
	// go allocate(noOfJobs)
	// done := make(chan bool)
	// go result(done)
	// noOfWorkers := 20
	// createWorkerPool(noOfWorkers)
	// <-done
	// endTime := time.Now()
	// diff := endTime.Sub(startTime)
	// fmt.Println("total time taken ", diff.Seconds(), "seconds")

	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)

	for {
		select {
		case s1, ok := <-output1:
			if !ok {
				fmt.Println("Server1 channel closed")
				output1 = nil
			} else {
				fmt.Println(s1)
			}
		case s2, ok := <-output2:
			if !ok {
				fmt.Println("Server2 channel closed")
				output2 = nil
			} else {
				fmt.Println(s2)
			}
		}
		if output1 == nil && output2 == nil {
			break
		}
	}
}
func server1(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server1"
	close(ch)
}
func server2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "from server2"
	close(ch)
}
	

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}


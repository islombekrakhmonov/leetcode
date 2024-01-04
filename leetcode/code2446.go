package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(haveConflict([]string{"01:15","02:00"},[]string{"02:00","03:00"}))
}

func haveConflict(event1 []string, event2 []string) bool {
	startEvent1,_ := time.Parse("15:04",event1[0])
	closeEvent1,_ := time.Parse("15:04",event1[1])
	startEvent2,_ := time.Parse("15:04",event2[0])
	closeEvent2,_ := time.Parse("15:04",event2[1])

	return !(startEvent1.After(closeEvent2) || startEvent2.After(closeEvent1))    
}
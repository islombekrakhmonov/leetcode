package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(daysBetweenDates("2019-01-01", "2020-01-01"))
}

func daysBetweenDates(date1 string, date2 string) int {
	t1, _ := time.Parse("2006-01-02", date1)
	t2, _ := time.Parse("2006-01-02", date2)

	return int(math.Abs(t1.Sub(t2).Hours() / 24))
}

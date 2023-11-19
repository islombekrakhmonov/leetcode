package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(dayOfYear("2019-01-01"))
}

func dayOfYear(date string) int {
	t, _ := time.Parse("2006-01-02", date)

	return t.YearDay()
}

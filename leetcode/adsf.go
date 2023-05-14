package main

import (
    "fmt"
    "time"
)

func main() {
    today := time.Now().Format("20060102")
    fmt.Println(today)
}

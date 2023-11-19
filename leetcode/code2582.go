package main

import "fmt"

func main() {
	fmt.Println(passThePillow(4, 5))
}

func passThePillow(n int, time int) int {
	var reverse bool
	var times int
	time++
	for time != 0 {
		if !reverse {
			if times == 0 {
				for i:=1; i<=n; i++ {
					time--
					if i == n {
						reverse = true
					}
					if time == 0 {
						return i
					}
				} 
			} else {
				for i:=2; i<=n; i++ {
					time--
					if i == n {
						reverse = true
					}
					if time == 0 {
						return i
					}
				} 
			}
		} else {
			for i:=n-1; i>0; i-- {
				time--
				if i == 1 {
					times++ 
					reverse = true
				}
				if time == 0 {
					return i
				}
			} 
		}
	}

	return time
}
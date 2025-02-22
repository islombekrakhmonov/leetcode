package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLatestTime("1?:?4")) //

}


//"?1:?6"
// "??:??"
// "??:1?"
// "?9:5?"
// "?0:40"

func findLatestTime(s string) string {

	for i := 0; i < len(s); i++ {
		switch i {
		case 0:
			if s[i] == '?' {

			}
		case 1:
			if s[i] == '?' {
			}
		case 2:
			if s[i] == '?' {
			}
		case 3:
			if s[i] == '?' {
			}
		}
	}

	return s
}

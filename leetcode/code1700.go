package main

import "fmt"

func main() {
	fmt.Println(countStudents([]int{1,1,1,0,0,1}, []int{1,0,0,0,1,1}))
	fmt.Println(countStudents1([]int{1,1,1,0,0,1}, []int{1,0,0,0,1,1}))
	fmt.Println(countStudents2([]int{1,1,1,0,0,1}, []int{1,0,0,0,1,1}))
}

func countStudents(students []int, sandwiches []int) int {

	found := true
	for found {
		found = false
		for i:=0; i<len(students); i++{
			if students[i] != sandwiches[i] {
				students = students[1:]
				students = append(students, students[i])
				found = true
				break
			} else {
				students = students[1:]
				sandwiches = sandwiches[:len(sandwiches)-1]
			}
		}
	}

	for len(students) > 0 {
        if students[0] == sandwiches[0] {
            students = students[1:]
            sandwiches = sandwiches[1:]
        } else {
            students = append(students, students[0])
            students = students[1:]
        }
    }

	return len(students)
}

func countStudents1(students []int, sandwiches []int) int {
    for len(students) > 0 {
        unableToEat := 0 
        n := len(students)
        
        for i := 0; i < n; i++ {
            if students[0] == sandwiches[0] {
                students = students[1:]
                sandwiches = sandwiches[1:]
                unableToEat = 0 
            } else {
                students = append(students, students[0])
                students = students[1:]
                unableToEat++
            }
        }
        
        if unableToEat == n {
            break
        }
    }
    
    return len(students)
}

func countStudents2(students []int, sandwiches []int) int {
    count := map[int]int{
        0: 0, 
        1: 0, 
    }

	
    for _, student := range students {
		count[student]++
    }
	
	fmt.Println(count)
    for _, sandwich := range sandwiches {
        if count[sandwich] > 0 {
            count[sandwich]--
        } else {
            break
        }
    }

    unableToEat := count[0] + count[1]

    return unableToEat
}
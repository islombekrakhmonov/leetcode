package main

import "fmt"

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
}

func isLongPressedName(name string, typed string) bool {

	if name[0] != typed[0] {
		return false
	}

	// lengthName, lengt hTyped := len(name), len(typed)
	prevChar, currChar := name[0], name[0]
	counter := 0

	for i := 0; i < len(typed); i++ {
		if typed[i] == currChar {
			counter++
		} else if typed[i] == prevChar {
			continue
		}

		if i == len(typed)-1 {
			if typed[i] != currChar && typed[i] != prevChar {
				fmt.Println(string(typed[i]), string(currChar), string(prevChar))
				return false
			}
		}
		prevChar, currChar = currChar, typed[i]
	}

	return true
}

// Remember that moving the index from Name, assign to prevChar = currChar and  currChar the new letter. Also check that the length is already reached during the loop.

// Check: name[0] == typed[0], if not -> false.
// make type variables: lengthName, lengthTyped.
// make char variables with index 0 from Name. char prevChar = Name[0], currChar = Name[0];
// make counter how many letters from Name match (index)
// loop where you check if letter from Typed is letter from currChar. If so, increment the counter (index name). If the letter does not agree, check if it agrees with the previous one (prevChar), If yes, do nothing.
// In the last subsection if the letter (from Typed) does not agree with CurrChar (name) and prevChar (name) return false.

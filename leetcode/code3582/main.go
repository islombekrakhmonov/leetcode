package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(generateTag("FkVsgzfhQxPYKr WtaCvDuHQeo roeVHIoPorZuIuMDxhYVs jlfCrXCDqBVjgmzxxmZOpUpCc  eHMpZrNxilBPngylMcS"))
}

func generateTag(caption string) string {

	var result = "#"

	caption = strings.TrimSpace(caption)

	if len(caption) == 0 {
		return result
	}

	splitted := strings.Split(caption, " ")
	for i := 0; i < len(splitted); i++ {
		if len(splitted[i]) == 0 {
			continue
		}
		if i == 0 {
			splitted[i] = strings.ToLower(string(splitted[i][0])) + strings.ToLower(splitted[i][1:])
		} else {
			splitted[i] = strings.ToUpper(string(splitted[i][0])) + strings.ToLower(splitted[i][1:])
		}

		for j := 0; j < len(splitted[i]); j++ {
			if (splitted[i][j] < 'A' || splitted[i][j] > 'Z') && (splitted[i][j] < 'a' || splitted[i][j] > 'z') {
				splitted[i] = splitted[i][:j] + splitted[i][j+1:]
			}
		}

		result += splitted[i]
	}

	if len(result) > 100 {
		return result[:100]
	}

	return result
}

/*
You are given a string caption representing the caption for a video.

The following actions must be performed in order to generate a valid tag for the video:

Combine all words in the string into a single camelCase string prefixed with '#'. A camelCase string is one where the first letter of all words except the first one is capitalized. All characters after the first character in each word must be lowercase.

Remove all characters that are not an English letter, except the first '#'.

Truncate the result to a maximum of 100 characters.

Return the tag after performing the actions on caption.

Note: The tag should be generated in lowercase.
*/

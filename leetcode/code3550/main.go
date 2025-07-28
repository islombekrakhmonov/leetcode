package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println(smallestIndex([]int{1, 2, 3, 4, 5}))
	// data := "ewogICJvcmRlcnMiOiBbCiAgICB7CiAgICAgICJzZXJ2aWNlSWQiOiAiMWE3OTI1ODgtMTc5Yi1jMzZmLTZiMGYtZWQzNGE3NDI0OGM0IiwKICAgICAgImZsaWdodE51bWJlciI6ICJLQyAxMTAiLAogICAgICAic2VydmljZURhdGUiOiAiMjAyNS0wNS0yNVQxMzozNTowMCIsCiAgICAgICJhZHVsdENvdW50IjogMSwKICAgICAgImNoaWxkQ291bnQiOiAwLAogICAgICAiaW5mYW50Q291bnQiOiAwLAogICAgICAiZmlyc3ROYW1lIjogIkZhcmtoYWQiLAogICAgICAibGFzdE5hbWUiOiAiQWJkdW1hemhpdHVseSIsCiAgICAgICAiZW1haWwiOiAiZi5hYmR1bWF6aGl0dWx5QGdtYWlsLmNvbSIsCiAgICAgICAicGhvbmUiOiIrNzcwNzAwMDAwNzciLAogICAgICAiY3VsdHVyZSI6ICJlbi1VUyIsCiAgICAgICJpc0RhdGVUaW1lT2ZQYXNzZW5nZXJzQXJyaXZhbFRvQWlycG9ydCI6IHRydWUKICAgIH0KICBdLAogICJwdWJsaWNBcGlLZXkiOiAiN2I3ZDcyNGUtZjgyNy03NTRjLTc0YmQtNTQ0OWVmNGViODJmIgp9"
	sign_string := "trE0Yzs8iCnm1F9mtRhn94L0bt17CfOjRcpL9pli" + "ewogICJvcmRlcnMiOiBbCiAgICB7CiAgICAgICJzZXJ2aWNlSWQiOiAiMWE3OTI1ODgtMTc5Yi1jMzZmLTZiMGYtZWQzNGE3NDI0OGM0IiwKICAgICAgImZsaWdodE51bWJlciI6ICJLQyAxMTAiLAogICAgICAic2VydmljZURhdGUiOiAiMjAyNS0xMC0zMCIsCiAgICAgICJhZHVsdENvdW50IjogMiwKICAgICAgImNoaWxkQ291bnQiOiAwLAogICAgICAiaW5mYW50Q291bnQiOiAwLAogICAgICAiZmlyc3ROYW1lIjogIkZhcmtoYWQiLAogICAgICAibGFzdE5hbWUiOiAiQWJkdW1hemhpdHVseSIsCiAgICAgICAiZW1haWwiOiAiaXNsb21iZWsucmFraG1vbm92QGVhc3l0by50cmF2ZWwiLAogICAgICAgInBob25lIjoiKzc3MDcwMDAwMDc3IiwKICAgICAgImN1bHR1cmUiOiAiZW4tVVMiLAogICAgICAib3RoZXJQYXNzZW5nZXJzQ29udGFjdERldGFpbHMiOiAiSmlsbCBEb2UiLAogICAgICAiaXNEYXRlVGltZU9mUGFzc2VuZ2Vyc0Fycml2YWxUb0FpcnBvcnQiOiB0cnVlCiAgICB9CiAgXSwKICAicHVibGljQXBpS2V5IjogIjdiN2Q3MjRlLWY4MjctNzU0Yy03NGJkLTU0NDllZjRlYjgyZiIKfQ==" + "trE0Yzs8iCnm1F9mtRhn94L0bt17CfOjRcpL9pli"
	signature := sha1ToBase64(sign_string)
	fmt.Println(signature)

	// departureTime := "2025-06-27T17:50:00Z"
	// departureTime1 := "2025-06-27T17:50:00"

	// var STANDART_TIME_FORMAT = "2006-01-02T15:04:05"

	// parsedDepartureTime, err := time.Parse(STANDART_TIME_FORMAT, departureTime)
	// fmt.Println(parsedDepartureTime, err)
	// parsedDepartureTime1, err := time.Parse(STANDART_TIME_FORMAT, departureTime1)
	// fmt.Println(parsedDepartureTime1, err)

	// fmt.Println(departureTime[:len(departureTime)-1])
	// // departureTimeParsed := time.Now()
	// departureTimeParsed, _ = time.Parse(STANDART_TIME_FORMAT, departureTime)

	//readJson

	// file, err := os.Open("./code3550/response.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// byteValue, err := io.ReadAll(file)
	// if err != nil {
	// 	panic(err)
	// }

	// var result map[string]interface{}
	// err = json.Unmarshal(byteValue, &result)
	// if err != nil {
	// 	panic(err)
	// }

	// products := result["services"].([]interface{})

	// for _, product := range products {
	// 	productMap := product.(map[string]interface{})
	// 	category := productMap["category"].(string)
	// 	serviceDirection := productMap["serviceDirection"].(string)
	// 	preOrderPeriodInMinutes := productMap["preOrderPeriodInMinutes"].(float64)
	// 	id := productMap["id"].(string)
	// 	if category == "FastTrack" && serviceDirection == "Arrival" {
	// 		if preOrderPeriodInMinutes > 0 {
	// 			fmt.Println(id)
	// 		}
	// 	}
	// }

}

func sha1ToBase64(data string) string {
	hash := sha1.Sum([]byte(data))                    // SHA1 hash (returns [20]byte)
	return base64.StdEncoding.EncodeToString(hash[:]) // Convert to base64 string
}

func smallestIndex(nums []int) int {
	for i, num := range nums {
		if i == sumOfDigits(num) {
			return i
		}
	}
	return -1
}

func sumOfDigits(num int) int {
	numsStr := fmt.Sprintf("%d", num)
	sum := 0
	for _, digit := range numsStr {
		sum += int(digit - '0')
	}
	return sum
}

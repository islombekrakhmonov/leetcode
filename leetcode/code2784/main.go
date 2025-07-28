package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
)

func main() {
	secretKey := []byte("AfmHsJfTBQLSu2D3ntaN7JHcARkJutiemq22bSISQZhYkDOw7sywF3x9Z2noxsTX")
	// simpleEncoding(secretKey)
	simpleEncodingRefund(secretKey)
}
func isGood(nums []int) bool {
	var (
		maxInt   = 0
		countMap = make(map[int]int)
	)
	for _, num := range nums {
		if num > maxInt {
			maxInt = num
		}
		countMap[num]++
	}

	if len(nums) != maxInt+1 {
		return false
	}

	for i := 1; i <= maxInt; i++ {
		if i == maxInt+1 {
			i = maxInt
		}
		if _, ok := countMap[i]; ok {
			countMap[i]--
			if countMap[i] == 0 {
				delete(countMap, i)
			}
		} else {
			return false
		}
	}

	if _, ok := countMap[maxInt]; ok {
		countMap[maxInt]--
		if countMap[maxInt] == 0 {
			delete(countMap, maxInt)
		}
	} else {
		return false
	}

	if len(countMap) != 0 {
		return false
	}

	return true
}

func isGood1(nums []int) bool {
	var (
		maxInt = 0
		array  []int
	)
	for _, num := range nums {
		if num > maxInt {
			maxInt = num
		}
	}

	for i := 1; i < maxInt+1; i++ {
		array = append(array, i)
	}

	array = append(array, maxInt)

	sort.Ints(nums)

	if len(array) != len(nums) {
		return false
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != array[i] {
			return false
		}
	}

	return true
}

func simpleEncodingRefund(secretKey []byte) {
	jsonData := []byte(`{"order_id":"24f4ef9e-ce07-4051-9c46-1aa6eb09ffa3","amount":129800,"currency":"UZS"}`)
	// jsonData := []byte(`{"psp_id":3623,"amount":42832400,"lifetime":5,"currency":"RUB","success_url":"https:\/\/easy-922-widget.easyto.travel\/W-7vd094z2yugj\/DXB\/P-000001185\/PV-000000593\/","fail_url":"https:\/\/test-widget.easyto.travel\/W-7vd094z2yugj\/vouchers","callback_url":"https:\/\/dev-test-api.easyto.travel\/ets_hub_webhooks\/test","details":{"order_id":"123","customer":"test user","email":"test@test.test"}}`)
	h := hmac.New(sha256.New, secretKey)
	h.Write(jsonData)
	hmacValue := h.Sum(nil)
	signature := hex.EncodeToString(hmacValue)
	fmt.Println("HMAC-SHA256 refund signature:", signature)
}

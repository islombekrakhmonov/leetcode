package main

import (
	"fmt"
	"regexp"
	"slices"
	"sort"
)

func main() {
	fmt.Println(validateCoupons([]string{"SAVE20", "", "PHARMA5", "SAVE@20"}, []string{"restaurant", "grocery", "pharmacy", "restaurant"}, []bool{true, true, true, true}))
}

type Coupon struct {
	Code         string
	BusinessLine string
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	var filtered []Coupon
	var categories = []string{"electronics", "grocery", "pharmacy", "restaurant"}
	categoryOrder := make(map[string]int)
	for i, cat := range categories {
		categoryOrder[cat] = i
	}

	for i := range code {
		if isValid(code[i]) && isActive[i] && contains(categories, businessLine[i]) {
			filtered = append(filtered, Coupon{
				Code:         code[i],
				BusinessLine: businessLine[i],
			})
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		if categoryOrder[filtered[i].BusinessLine] != categoryOrder[filtered[j].BusinessLine] {
			return categoryOrder[filtered[i].BusinessLine] < categoryOrder[filtered[j].BusinessLine]
		}
		return filtered[i].Code < filtered[j].Code
	})

	var output []string
	for _, c := range filtered {
		output = append(output, c.Code)
	}

	return output
}

func isValid(coupon string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

	return re.MatchString(coupon) && len(coupon) > 0
}

func contains(slice []string, element string) bool {
	return slices.Contains(slice, element)
}

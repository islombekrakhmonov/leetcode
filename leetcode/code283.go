package main

func main() {
	moveZeroes([]int{0, 4, 0, 3, 12})
}

func moveZeroes(nums []int) {

	nonZeroIndex := 0

	n := len(nums)

	for i := 1; i < n; i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex], nums[i] = nums[i], nums[nonZeroIndex]
			nonZeroIndex++
		}
	}

	return
}

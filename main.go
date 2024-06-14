package main

import "fmt"

func main() {
	input := []int{2, 1, -1}
	output := pivotIndex(input)
	fmt.Println(output)
}

func pivotIndex(nums []int) int {
	sum, leftSum := 0, 0
	for _, v := range nums {
		sum += v
	}

	for i := 0; i < len(nums); i++ {
		if leftSum == sum-nums[i]-leftSum {
			return i
		}

		leftSum += nums[i]
	}

	return -1
}

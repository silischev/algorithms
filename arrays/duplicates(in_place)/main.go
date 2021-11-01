package main

import "log"

func main() {
	nums := []int{1}
	//nums := []int{1, 1}
	//nums := []int{1, 1, 2}
	//nums := []int{1, 2, 3}
	//nums := []int{1, 1, 1, 2}
	//nums := []int{1, 2, 2, 3}
	log.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) > 30000 || len(nums) == 0 {
		return 0
	}

	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] || nums[i] < -100 || nums[i] > 100 {
			nums = append(nums[:i+1], nums[i+2:]...)
			continue
		}

		i++
	}

	log.Println(nums)

	return len(nums)
}

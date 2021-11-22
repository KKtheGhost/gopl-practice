package main

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("no min for empty slice")
	}
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("no max for empty slice")
	}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func min2(base int, nums ...int) int {
	min := base
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}

func max2(base int, nums ...int) int {
	min := base
	for _, n := range nums {
		if n > min {
			min = n
		}
	}
	return min
}

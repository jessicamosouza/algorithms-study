package binarysearch

func BinarySearch(nums []int, target int) int {
	var min int
	max := len(nums) - 1

	for i := range nums {
		av := (max + min) / 2

		if target > nums[av] {
			min = i - 1
		}
		if target < nums[av] {
			max = i + 1
		}
		if nums[i] == target {
			return i
		}
	}
	return -1
}

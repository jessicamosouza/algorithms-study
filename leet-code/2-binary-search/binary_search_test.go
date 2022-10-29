package binarysearch_test

import (
	"testing"

	binarysearch "problems/leet-code/2-binary-search"
)

// https://leetcode.com/problems/binary-search/

// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4
// Example 2:

// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1

// Constraints:

// 1 <= nums.length <= 104
// -104 < nums[i], target < 104
// All the integers in nums are unique.
// nums is sorted in ascending order.

func TestBinarySearch(t *testing.T) {
	testCase := []struct {
		nums   []int
		target int
		output int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			output: 4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			output: -1,
		},
	}

	for _, test := range testCase {
		result := binarysearch.BinarySearch(test.nums, test.target)
		if result != test.output {
			t.Errorf("Função retornou %d, porém era para retornar %d, dado os valores de %v e %d", result, test.output, test.nums, test.target)
		}
	}

}

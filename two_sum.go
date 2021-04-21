/*
Two Sum Problem
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:
2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/

package main

import (
	"fmt"
)

type Tests struct {
	nums   []int
	target int
}

func main() {
	tests := []Tests{
		{nums: []int{2, 7, 11, 15}, target: 9},
		{nums: []int{3, 2, 4}, target: 6},
		{nums: []int{3, 3}, target: 6},
		{nums: []int{1, 6, 20, 15, -1, 9}, target: 8},
		{nums: []int{}, target: 0},
	}
	for _, test := range tests {
		fmt.Printf("Brute Force: %v\n", BruteForce(test.nums, test.target))
		fmt.Printf("Hashing: %v\n", Hashing(test.nums, test.target))
	}
}

func BruteForce(nums []int, target int) []int {
	// Time Complexity: O(n^2)
	// Space Complexity: N/A
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func Hashing(nums []int, target int) []int {
	// Time Complexity O(n)
	// Space Complexity O(n)
	set := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := set[target-nums[i]]; ok {
			return []int{set[target-nums[i]], i}
		}
		set[nums[i]] = i
	}
	return []int{}
}

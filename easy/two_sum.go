package easy

/*
1. Two Sum (leetcode.com/problems/two-sum)

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:
	Input: nums = [2,7,11,15], target = 9
	Output: [0,1]
	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
	Input: nums = [3,2,4], target = 6
	Output: [1,2]

Example 3:
	Input: nums = [3,3], target = 6
	Output: [0,1]



Constraints:
  2 <= nums.length <= 104
  -109 <= nums[i] <= 109
  -109 <= target <= 109
  Only one valid answer exists.



Follow-up: Can you come up with an algorithm that is less than O(n^2) time complexity?
*/

func betterTwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	// nums = [3,2,4], target 6

	for i, v := range nums {
		numMap[v] = i
		// 3 -> 0, 2 -> 1, 4 ->
	}

	for i, v := range nums {
		diff := target - v

		// val is the value off diff from the map if it exists
		// ok is a bool that shows if diff exists in the map
		if val, ok := numMap[diff]; ok && (val != i) {
			return []int{i, val}
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
	var result []int

	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			// k = i + 1, do not sum with self
			if v+nums[k] == target {
				result = append(result, i, k)
			}
		}
	}

	return result
}

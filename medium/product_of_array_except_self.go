package medium

/*
238. Product of Array Except Self (leetcode.com/problems/product-of-array-except-self)

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:
	Input: nums = [1,2,3,4]
	Output: [24,12,8,6]

Example 2:
	Input: nums = [-1,1,0,-3,3]
	Output: [0,0,9,0,0]



Constraints:
    2 <= nums.length <= 10^5
    -30 <= nums[i] <= 30
    The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.



Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/

func productExceptSelf(nums []int) []int { // Fails, TLE
	var result []int

	for i := range nums {
		product := 1

		for k, num := range nums {
			if i != k {
				product *= num
			}
		}

		result = append(result, product)
	}

	return result
}

func productExceptSelf2(nums []int) []int {
	size := len(nums)
	result := make([]int, size)
	result[0], result[size-1] = 1, 1

	for i := 1; i < size; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := size - 2; i >= 0; i-- {
		rightProduct *= nums[i+1]
		result[i] *= rightProduct
	}

	return result
}

func productExceptSelf3(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = 1
	}

	pref := 1
	for i := 0; i < len(nums); i++ {
		result[i] *= pref
		pref *= nums[i]
	}

	post := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= post
		post *= nums[i]
	}

	return result
}

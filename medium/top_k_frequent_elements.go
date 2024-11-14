package medium

import "sort"

/*
347. Top K Frequent Elements (leetcode.com/problems/top-k-frequent-elements)

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.



Example 1:
	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]

Example 2:
	Input: nums = [1], k = 1
	Output: [1]



Constraints:
    1 <= nums.length <= 10^5
    -10^4 <= nums[i] <= 10^4
    k is in the range [1, the number of unique elements in the array].
    It is guaranteed that the answer is unique.
*/

type kv struct {
	key int
	val int
}

func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	var result []int
	var ss []kv

	for _, num := range nums {
		numMap[num]++
	}

	for k, v := range numMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	for _, ns := range ss[:k] {
		result = append(result, ns.key)
	}

	return result
}

func topKFrequent2(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	uniques := make([]int, 0, len(numMap))
	for num := range numMap {
		uniques = append(uniques, num)
	}

	sort.Slice(uniques, func(i, j int) bool {
		return numMap[uniques[i]] > numMap[uniques[j]]
	})

	return uniques[:k]
}

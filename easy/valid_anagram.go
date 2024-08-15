package easy

/*
242. Valid Anagram (leetcode.com/problems/valid-anagram)

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:
	Input: s = "anagram", t = "nagaram"
	Output: true

Example 2:
	Input: s = "rat", t = "car"
	Output: false



Constraints:
  1 <= s.length, t.length <= 5 * 10^4
  s and t consist of lowercase English letters.



Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

func isAnagram(s string, t string) bool {
	chars := make(map[rune]int)

	for _, l := range s {
		chars[l]++
	}

	for _, l := range t {
		chars[l]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}

package easy

/*
20. Valid Parentheses (leetcode.com/problems/valid-parentheses)

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
  Open brackets must be closed by the same type of brackets.
  Open brackets must be closed in the correct order.
  Every close bracket has a corresponding open bracket of the same type.



Example 1:
	Input: s = "()"
	Output: true

Example 2:
	Input: s = "()[]{}"
	Output: true

Example 3:
	Input: s = "(]"
	Output: false



Constraints:
  1 <= s.length <= 104
  s consists of parentheses only '()[]{}'.
*/

var pMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	var stack []rune

	for _, char := range s {
		if close, ok := pMap[char]; ok {
			stack = append(stack, close)
		} else if len(stack) == 0 || stack[len(stack)-1] != char {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

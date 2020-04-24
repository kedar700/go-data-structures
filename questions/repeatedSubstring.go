package questions

import (
	"strings"
)

/**
Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together. You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.



Example 1:

Input: "abab"
Output: True
Explanation: It's the substring "ab" twice.
Example 2:

Input: "aba"
Output: False
Example 3:

Input: "abcabcabcabc"
Output: True
Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)
if S is composed of k substrings s, then S1 = S + S should contain 2k substrings. Destroying the first and the last character leads to at least (2k - 2) substrings left.

since k >= 2,
2k - 2 >= k
which means that S1[1:-1] should still contain S
*/
func RepeatedSubstringPattern(s string) bool {
	patternToCheck := (s + s)[1 : (len(s)*2)-1]
	return strings.Contains(patternToCheck, s)
}

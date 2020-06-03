package medium

import (
	"fmt"
	"regexp"
	"strings"
)

/**

Given an input string, reverse the string word by word.

Example 1:

Input: "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: "  hello world!  "
Output: "world! hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


Note:

A word is defined as a sequence of non-space characters.
Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
You need to reduce multiple spaces between two words to a single space in the reversed string.

*/

func reverseWords(s string) string {
	fmt.Println("reverse the string word by word")

	space := regexp.MustCompile(`\s+`)
	s = space.ReplaceAllString(strings.TrimSpace(s), " ")

	strA := strings.Split(s, " ")
	l := len(strA)

	for i := 0; i < l/2; i++ {
		strA[i], strA[l-i-1] = strings.TrimSpace(strA[l-i-1]), strings.TrimSpace(strA[i])
	}

	output := strings.Join(strA, " ")

	return output
}

// DoReverseWords function to reverse the words in a string
func DoReverseWords() {
	// str := "SOme thing is not going right"
	str := "  hello world!  "

	fmt.Println(reverseWords(str))

}

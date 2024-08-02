package main

import "unicode"

/*
@Time : 2024/7/6 20:37
@Author : echo
@File : word2
@Software: GoLand
@Description:
*/
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

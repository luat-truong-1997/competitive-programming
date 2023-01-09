/**
* Problem: https://leetcode.com/problems/letter-case-permutation/
* Generate permutation by recursive
 */
package main

import "fmt"

func letterCasePermutation(s string) []string {
	r := []rune(s)
	res := make([]string, 0)
	return generate(r, 0, res, "")
}

func generate(r []rune, index int, res []string, current string) []string {
	if index == len(r) {
		return append(res, current)
	}
	newCurrent := current + string(r[index])
	res = generate(r, index+1, res, newCurrent)
	if isLower(r[index]) || isUpper(r[index]) {
		newCurrent := current + string(transform(r[index]))
		res = generate(r, index+1, res, newCurrent)
	}
	return res
}

func isLower(s rune) bool {
	return 'a' <= s && s <= 'z'
}

func isUpper(s rune) bool {
	return 'A' <= s && s <= 'Z'
}

func transform(s rune) rune {
	if isLower(s) {
		return 'A' + s - 'a'
	}
	if isUpper(s) {
		return 'a' + s - 'A'
	}
	return s
}

func main() {
	s := "a1b2"
	fmt.Println(letterCasePermutation(s))
}

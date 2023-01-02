/**
* Problem : https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings
* Implement recursion. Time complexity is sum of every combination of n => 2^n
 */
package main

func maxUniqueSplit(s string) int {
	r := []rune(s)
	result := make([]string, 0)
	has := make(map[string]bool)
	maxResult := 0
	splitUnique(0, r, "", has, result, &maxResult)
	return maxResult
}

func splitUnique(index int, r []rune, current string, has map[string]bool, result []string, maxResult *int) {
	if index == len(r) {
		if *maxResult < len(result) {
			*maxResult = len(result)
		}
		return
	}
	current = current + string(r[index])
	if _, ok := has[current]; !ok {
		has[current] = true
		result = append(result, current)
		splitUnique(index+1, r, "", has, result, maxResult)

		// Revert recurrsion
		delete(has, current)
		result = result[0 : len(result)-1]
	}
	splitUnique(index+1, r, current, has, result, maxResult)
}

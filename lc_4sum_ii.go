/**
* Problem : https://leetcode.com/problems/4sum-ii
* Build sum array from each 2 array. Complexity : n^2
* Sort 1 array then binary search on it
 */
package main

import "sort"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum12 := buildSum(nums1, nums2)
	sum34 := buildSum(nums3, nums4)
	result := 0
	sort.Ints(sum34)
	for _, item := range sum12 {
		leftIndex := sort.SearchInts(sum34, -item)
		if leftIndex == len(sum34) || sum34[leftIndex] != -item {
			continue
		}
		rightIndex := sort.SearchInts(sum34, -item+1)
		result += rightIndex - leftIndex
	}
	return result
}

func buildSum(a, b []int) []int {
	result := make([]int, 0)
	for _, item := range a {
		for _, item2 := range b {
			result = append(result, item+item2)
		}
	}
	return result
}

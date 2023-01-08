/**
* Problem: https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array
* Use Trie. Prioritize the reversed bit.
 */
package main

import "fmt"

type trie struct {
	children map[int]*trie
	value    int
}

func newTrie() *trie {
	children := make(map[int]*trie)
	return &trie{
		children: children,
		value:    0,
	}
}

func (t *trie) add(index int, nums []int, value int) {
	if index == len(nums) {
		t.value = value
		return
	}
	child, ok := t.children[nums[index]]
	if !ok {
		child = newTrie()
		t.children[nums[index]] = child
	}
	child.add(index+1, nums, value)

}

func (t *trie) query(index int, nums []int, value int) int {
	if index == len(nums) {
		return value ^ t.value
	}
	x := 1 - nums[index]
	child, ok := t.children[x]
	if ok {
		return child.query(index+1, nums, value)
	}
	child, ok = t.children[nums[index]]
	if ok {
		return child.query(index+1, nums, value)
	}
	return 0
}

func numberToBit(nums int) []int {
	result := make([]int, 33)
	for i := 32; i >= 0; i-- {
		result[32-i] = (nums >> i) & 1
	}
	return result
}

func findMaximumXOR(nums []int) int {
	root := newTrie()
	result := 0
	for _, num := range nums {
		bits := numberToBit(num)
		root.add(0, bits, num)
		q := root.query(0, bits, num)
		if result < q {
			result = q
		}
	}

	return result
}

func main() {
	arr := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR(arr))
}

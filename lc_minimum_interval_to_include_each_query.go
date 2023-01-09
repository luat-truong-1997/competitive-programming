/**
* Problem : https://leetcode.com/problems/minimum-interval-to-include-each-query
* Optimize from brute forces (find min(interval[i][r]-interval[i][l]+1)
* with condition interval[i][l] <= query[x] < interval[i][r])    solution.
* First, let sort interval in asc order by left and sort query in asc order.
* So, I will can clear the condition interval[i][l] <= query by pointer technique.
* For the condition query[x] < interval[i][r], I'll use segment tree.
 */
package main

import (
	"sort"
)

type segmentTree struct {
	l, r                  int
	leftChild, rightChild *segmentTree
	value                 int
}

func newSegmentTreeNode(l, r int) *segmentTree {
	return &segmentTree{
		l:          l,
		r:          r,
		leftChild:  nil,
		rightChild: nil,
		value:      -1,
	}
}

func (node *segmentTree) addValue(index, value int) {
	if node.l == index && node.r == index {
		node.value = min(node.value, value)
		return
	}
	if node.l > index {
		return
	}
	if node.r < index {
		return
	}
	m := (node.l + node.r) / 2
	if node.leftChild == nil {
		node.leftChild = newSegmentTreeNode(node.l, m)
	}
	if node.rightChild == nil {
		node.rightChild = newSegmentTreeNode(m+1, node.r)
	}
	node.leftChild.addValue(index, value)
	node.rightChild.addValue(index, value)
	node.value = min(node.leftChild.value, node.rightChild.value)

}

func (node *segmentTree) query(l, r int) int {
	if node == nil {
		return -1
	}
	if l <= node.l && node.r <= r {
		return node.value
	}
	if node.l > r || node.r < l {
		return -1
	}
	return min(node.leftChild.query(l, r), node.rightChild.query(l, r))
}

func min(a, b int) int {
	if a == -1 {
		return b
	}
	if b == -1 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

type intervalSorters struct {
	data [][]int
	cmp  func(a, b []int) bool
}

func (s *intervalSorters) Len() int {
	return len(s.data)
}

func (s *intervalSorters) Less(i, j int) bool {
	return s.cmp(s.data[i], s.data[j])
}

func (s *intervalSorters) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

var compareLeft = func(a, b []int) bool {
	if a[0] < b[0] {
		return true
	}
	if a[0] > b[0] {
		return false
	}
	if a[1] < b[1] {
		return true
	}
	return false
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Sort(&intervalSorters{data: intervals, cmp: compareLeft})
	maxRight := 0
	for _, value := range intervals {
		if maxRight < value[1] {
			maxRight = value[1]
		}
	}
	queriesParsed := make([][]int, len(queries))
	for index, value := range queries {
		queriesParsed[index] = make([]int, 2)
		queriesParsed[index][0] = value
		queriesParsed[index][1] = index
	}
	sort.Sort(&intervalSorters{data: queriesParsed, cmp: compareLeft})
	results := make([]int, len(queries))
	interValPointer := 0
	segmentTree := newSegmentTreeNode(1, maxRight)
	for _, query := range queriesParsed {
		for interValPointer < len(intervals) && intervals[interValPointer][0] <= query[0] {
			segmentTree.addValue(intervals[interValPointer][1], intervals[interValPointer][1]-intervals[interValPointer][0]+1)
			interValPointer++
		}
		if query[0] <= maxRight {
			results[query[1]] = segmentTree.query(query[0], maxRight)
		} else {
			results[query[1]] = -1
		}
	}

	return results
}

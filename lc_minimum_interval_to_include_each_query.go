package main

import "sort"

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

var compareRight = func(a, b []int) bool {
	if a[1] < b[1] {
		return true
	}
	if a[1] > b[1] {
		return false
	}
	if a[0] < b[0] {
		return true
	}
	return false
}

func minInterval(intervals [][]int, queries []int) []int {

	sort.Sort(&intervalSorters{data: intervals, cmp: compareLeft})
	results := make([]int, 0)
	for _, query := range queries {
		result := -1
		index := sort.Search(len(intervals), func(i int) bool { return intervals[i][0] <= query })
		if index < len(intervals) {
			result = intervals[index][1] - intervals[index][0]
		}
		results = append(results, result)
	}
	sort.Sort(&intervalSorters{data: intervals, cmp: compareRight})
	for i, query := range queries {
		result := -1
		index := sort.Search(len(intervals), func(i int) bool { return intervals[i][1] <= query })
		if index < len(intervals) {
			result = intervals[index][1] - intervals[index][0]
		}
		if results[i] == -1 {
			results[i] = result
		}
		if result != -1 && results[i] > result {
			results[i] = result
		}
		results = append(results, result)
	}
	return results
}

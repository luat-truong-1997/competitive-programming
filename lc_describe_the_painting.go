/**
* Problem : https://leetcode.com/problems/describe-the-painting
* After sort point and color, use 2 pointers to make result
 */
package main

import (
	"fmt"
	"sort"
)

type colorSet struct {
	set map[int64]int
	sum int64
}

func (s *colorSet) addSet(set *colorSet) {
	for key, value := range set.set {
		s.add(key, value)
	}
}

func (s *colorSet) add(color int64, value int) {
	count, ok := s.set[color]
	count += value
	s.set[color] = count
	if !ok {
		s.sum += color
	}
}

func (s *colorSet) remove(color int64, cnt int) {
	count, _ := s.set[color]
	count -= cnt
	if count == 0 {
		delete(s.set, color)
		s.sum -= color
	} else {
		s.set[color] = count
	}
}

func (s *colorSet) removeSet(set *colorSet) {
	for key, value := range set.set {
		s.remove(key, value)
	}
}

func (s *colorSet) value() int64 {
	return s.sum
}

type timePoint struct {
	point  int
	colors *colorSet
}

func newColorSet() *colorSet {
	return &colorSet{
		set: make(map[int64]int),
		sum: 0,
	}
}
func newTimePoint(point int, color int64) *timePoint {
	colors := newColorSet()
	colors.add(color, 1)
	return &timePoint{
		point:  point,
		colors: colors,
	}
}

func splitPainting(segments [][]int) [][]int64 {
	result := make([][]int64, 0)
	leftTimePoints := make([]*timePoint, 0)
	rightTimePoints := make([]*timePoint, 0)
	leftTimeMap := make(map[int]*timePoint)
	rightTimeMap := make(map[int]*timePoint)
	currentSet := newColorSet()

	for _, segment := range segments {
		left, right, color := segment[0], segment[1], segment[2]
		leftPoints, ok := leftTimeMap[left]
		if !ok {
			leftPoints = newTimePoint(left, int64(color))
			leftTimeMap[left] = leftPoints
			leftTimePoints = append(leftTimePoints, leftPoints)
		} else {
			leftPoints.colors.add(int64(color), 1)
		}

		rightPoints, ok := rightTimeMap[right]
		if !ok {
			rightPoints = newTimePoint(right, int64(color))
			rightTimeMap[right] = rightPoints
			rightTimePoints = append(rightTimePoints, rightPoints)
		} else {
			rightPoints.colors.add(int64(color), 1)
		}
	}

	sort.Slice(leftTimePoints, func(i, j int) bool { return leftTimePoints[i].point < leftTimePoints[j].point })
	sort.Slice(rightTimePoints, func(i, j int) bool { return rightTimePoints[i].point < rightTimePoints[j].point })

	l := leftTimePoints[0].point
	currentSet = leftTimePoints[0].colors
	left, right := 1, 0
	for left < len(leftTimePoints) && right < len(rightTimePoints) {
		if leftTimePoints[left].point < rightTimePoints[right].point {
			currentColor := currentSet.value()
			currentSet.addSet(leftTimePoints[left].colors)
			nextColor := currentSet.value()
			if currentColor != nextColor && l != leftTimePoints[left].point {
				if currentColor != 0 {
					result = append(result, []int64{int64(l), int64(leftTimePoints[left].point), currentColor})
				}
				l = leftTimePoints[left].point
			}
			left++
		} else {
			currentColor := currentSet.value()
			currentSet.removeSet(rightTimePoints[right].colors)
			nextColor := currentSet.value()
			if currentColor != nextColor {
				if currentColor != 0 {

					result = append(result, []int64{int64(l), int64(rightTimePoints[right].point), currentColor})
				}
				l = rightTimePoints[right].point
			}
			right++
		}
	}
	for left < len(leftTimePoints) {
		currentColor := currentSet.value()
		currentSet.addSet(leftTimePoints[left].colors)
		nextColor := currentSet.value()
		if currentColor != nextColor && l != leftTimePoints[left].point {
			if currentColor != 0 {
				result = append(result, []int64{int64(l), int64(leftTimePoints[left].point), currentColor})
			}
			l = leftTimePoints[left].point
		}
		left++
	}
	for right < len(rightTimePoints) {
		currentColor := currentSet.value()
		currentSet.removeSet(rightTimePoints[right].colors)
		nextColor := currentSet.value()
		if currentColor != nextColor {
			if currentColor != 0 {
				result = append(result, []int64{int64(l), int64(rightTimePoints[right].point), currentColor})
			}
			l = rightTimePoints[right].point
		}
		right++
	}
	return result
}

func main() {
	segments := [][]int{{4, 5, 9},
		{8, 12, 5},
		{4, 7, 19},
		{14, 15, 1}, {3, 10, 8}, {17, 20, 18}, {7, 19, 14}, {8, 16, 6}, {14, 17, 7}, {11, 13, 3}}
	fmt.Println(splitPainting(segments))
}

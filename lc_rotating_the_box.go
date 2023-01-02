/**
* Problem : https://leetcode.com/problems/rotating-the-box
* Ad-hoc problem. Just do what they told
 */
package main

func rotateTheBox(box [][]byte) [][]byte {
	n := len(box)
	m := len(box[0])
	result := make([][]byte, m)
	for i := range result {
		result[i] = make([]byte, n)
	}
	for x := range box {
		for y := range box[x] {
			newX, newY := findItemPosition(x, y, n, m)
			result[newX][newY] = box[x][y]
		}
	}
	gravity(result)
	return result
}

func gravity(box [][]byte) {
	stone := byte('#')
	empty := byte('.')
	n := len(box)
	for y := range box[0] {
		for x := n - 1; x >= 0; x-- {
			if box[x][y] == stone {
				newX := x
				for newX+1 < n && box[newX+1][y] == empty {
					newX++
				}
				box[x][y] = empty
				box[newX][y] = stone
			}
		}
	}
}

func findItemPosition(x, y, n, m int) (int, int) {
	return findItemNewX(y, m), findItemNewY(x, n)
}

func findItemNewX(y, m int) int {
	return y
}

func findItemNewY(x, n int) int {
	return n - x - 1
}

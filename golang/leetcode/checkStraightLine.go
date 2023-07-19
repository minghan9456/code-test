package main

import "fmt"

func main() {
	// coordinates := [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	// coordinates := [][]int{{0, 1}, {0, -1}, {0, 4}}
	// coordinates := [][]int{{1, 1}, {0, 1}, {4, 1}, {0, 0}}
	coordinates := [][]int{{2, 1}, {4, 2}, {6, 3}}
	fmt.Println(checkStraightLine(coordinates))
}

func checkStraightLine(coordinates [][]int) bool {
	x1 := coordinates[0][0]
	y1 := coordinates[0][1]
	x2 := coordinates[1][0]
	y2 := coordinates[1][1]

	// y2 - y1 / x2 - x1 == yn - y1 / xn - x1

	for i := 2; i < len(coordinates); i++ {
		xn := coordinates[i][0]
		yn := coordinates[i][1]

		if (y2-y1)*(xn-x1) != (yn-y1)*(x2-x1) {
			return false
		}
	}

	return true
}

/*
func check(p1, p2 []int) (bool, bool, float64) {
	var m float64
	vertical := false
	horizontal := false

	if p2[0]-p1[0] == 0 {
		// x
		horizontal = true
	} else if p2[1]-p1[1] == 0 {
		// y
		vertical = true
	} else {
		m = float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
	}

	return vertical, horizontal, m
}

func checkStraightLine(coordinates [][]int) bool {
	p1 := coordinates[0]
	p2 := coordinates[1]

	vertical, horizontal, m := check(p1, p2)

	for i := 2; i < len(coordinates); i++ {
		pn := coordinates[i]
		nv, nh, nm := check(p1, pn)

		if nv != vertical || nh != horizontal || nm != m {
			return false
		}
	}

	return true
}
*/

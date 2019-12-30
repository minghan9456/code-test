package main

import "fmt"

func main() {

	var dishes [][]string
	dishes = [][]string{
		{"Salad", "Tomato", "Cucumber", "Salad", "Sauce"},
		{"Pizza", "Tomato", "Sausage", "Sauce", "Dough"},
		{"Quesadilla", "Chicken", "Cheese", "Sauce"},
		{"Sandwich", "Salad", "Bread", "Tomato", "Cheese"},
	}

	fmt.Println(groupingDishes(dishes))

}

func groupingDishes(dishes [][]string) [][]string {
	var ret [][]string
	var clct map[string][]string

	ret = make([][]string, 0)
	clct = make(map[string][]string)

	for i := 0; i <= len(dishes)-1; i++ {
		dish := dishes[i][0]
		for j := 1; j <= len(dishes[i])-1; j++ {
			ingredients := dishes[i][j]
			_, ok := clct[dishes[i][j]]
			if !ok {
				clct[ingredients] = []string{dish}
			} else {
				clct[ingredients] = append(clct[ingredients], dish)
			}
		}
	}

	fmt.Println(clct)

	for ingredients, dish := range clct {
		if len(dish) >= 2 {
			fmt.Println(ingredients)
			quickSort_1(dish, 0, len(dish)-1)
			ret = append(ret, append([]string{ingredients}, dish...))
		}
	}

	quickSort(ret, 0, len(ret)-1)

	return ret
}

func partition(a [][]string, lo, hi int) int {
	p := a[hi][0]
	for j := lo; j < hi; j++ {
		if a[j][0] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a [][]string, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func partition_1(a []string, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort_1(a []string, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition_1(a, lo, hi)
	quickSort_1(a, lo, p-1)
	quickSort_1(a, p+1, hi)
}

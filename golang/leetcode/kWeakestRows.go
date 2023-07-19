// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

func kWeakestRows(mat [][]int, k int) []int {
	clct := [][2]int{} // {sum, index}

	// Time: O(m*n)
	for i, row := range mat {
		sum := 0
		for _, col := range row {
			sum += col
		}
		clct = append(clct, [2]int{sum, i})
	}

	//[2 4 1 2 5]
	// Sort by sum & index
	// Time: O(nlogn)
	sort.Slice(clct, func(a, b int) bool {
		return (clct[a][0] < clct[b][0]) || (clct[a][0] == clct[b][0] && clct[a][1] < clct[b][1])
	})

	// Time: O(k)
	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, clct[i][1])
	}

	return ans
}

func kWeakestRows(mat [][]int, k int) []int {
	myRows := make(weakestRows, len(mat))

	for i, row := range mat {
		val := 0
		for idx := 0; idx < len(row); idx++ {
			val += row[idx]
		}

		myRows[i] = weakestRow{Val: val, Idx: i}
	}

	heap.Init(&myRows)
	res := []int{}

	for j := 0; j < k; j++ {
		row := heap.Pop(&myRows).(weakestRow)
		res = append(res, row.Idx)
	}

	return res
}

type weakestRow struct {
	Val int
	Idx int
}

type weakestRows []weakestRow

func (wrs weakestRows) Len() int {
	return len(wrs)
}

func (wrs weakestRows) Swap(i, j int) {
	wrs[i], wrs[j] = wrs[j], wrs[i]
}

func (wrs weakestRows) Less(i, j int) bool {
	if wrs[i].Val == wrs[j].Val {
		return wrs[i].Idx < wrs[j].Idx
	} else if wrs[i].Val > wrs[j].Val {
		return false
	} else {
		return true
	}
}

func (wrs *weakestRows) Pop() interface{} {
	l := len(*wrs) - 1
	res := (*wrs)[l]
	*wrs = (*wrs)[:l]
	return res
}

func (wrs *weakestRows) Push(newRow interface{}) {
	row := newRow.(weakestRow)
	*wrs = append(*wrs, row)
}

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

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

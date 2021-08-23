// https://leetcode.com/problems/all-paths-from-source-to-target/

func DFS (curr int, graph *[][]int, res *[][]int, clct []int) {
    
    clct = append(clct, curr)
        
    l := len((*graph)) - 1
    if curr == l {
        //fmt.Println(curr, clct, "add")
        // A slice is passed by reference, not value, so we must make a copy of the slice for each child.
        cp := make([]int, len(clct))
        copy(cp, clct)
        (*res) = append((*res), cp) 
    }

    for _, nei := range (*graph)[curr] {
        // fmt.Println("here", clct)
        DFS(nei, graph, res, clct)
    }
    
    return
}

func allPathsSourceTarget(graph [][]int) [][]int {
    res := [][]int{}
    
    clct := []int{}
    DFS(0, &graph, &res, clct)
    
    // fmt.Println(res)
    
    return res
}

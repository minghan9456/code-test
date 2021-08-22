// https://leetcode.com/problems/find-center-of-star-graph/
// Reading the description carefully helps here. Basically, a center node must appear in every edge.

func findCenter(edges [][]int) int {
    if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
        return edges[0][0]
    } else {
        return edges[0][1]
    }
}

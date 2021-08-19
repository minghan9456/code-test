// https://leetcode.com/problems/clone-graph

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// BFS
func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    
    m := map[*Node]*Node{node: nil}
    q := []*Node{node}
    
    m[node] = &Node{Val:node.Val}
    
    for len(q) > 0 {
        curr := q[0]
        q = q[1:]
        
        for i := range curr.Neighbors {
            neighbor := curr.Neighbors[i]
            
            if _, seen := m[neighbor]; !seen {
                m[neighbor] = &Node{Val: neighbor.Val}
                q = append(q, neighbor)
            }
            m[curr].Neighbors = append(m[curr].Neighbors, m[neighbor])
        }    
    }
    
    return m[node]
}

// DFS
func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    visited := make(map[int]*Node)
    return walk(node, visited)
}

func walk(node *Node, visited map[int]*Node) *Node {
    if n, ok := visited[node.Val]; ok {
        return n
    }
    n := &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
    visited[node.Val] = n
    for i, _ := range node.Neighbors {
        n.Neighbors[i] = walk(node.Neighbors[i], visited)
    }
    return n
}

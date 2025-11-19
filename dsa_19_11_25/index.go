package dsa_19_11_25

//https://leetcode.com/problems/clone-graph/description/

type Node struct {
	Val       int
	Neighbors []*Node
}

func dfs(node *Node, visited map[*Node]*Node) *Node {

	// if node is nil then 
    if node == nil {
        return nil
    }

	// if visited stop
    if v, ok := visited[node]; ok {
        return v
    }

	// clone the node
    cloneNode := &Node{
        Val: node.Val,
    }

	// add to visited
    visited[node] = cloneNode

	// handle the neigbours recursively
    for _, n := range node.Neighbors {
        cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(n, visited))
    }

    return cloneNode
}


func solution(node *Node) *Node {

	if node == nil {
        return nil
    }

	visited := make(map[*Node]*Node)
    return dfs(node, visited)
}

func Main() {

}

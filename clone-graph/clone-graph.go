package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	memo := make(map[*Node]*Node, 0)
	var dfs func(*Node) *Node
	dfs = func(cn *Node) *Node {
		if clone, exists := memo[cn]; exists {
			return clone
		}

		copiedN := &Node{Val: cn.Val, Neighbors: make([]*Node, 0, len(cn.Neighbors))}

		memo[cn] = copiedN

		for _, neigh := range cn.Neighbors {
			if clone, exists := memo[neigh]; exists {
				copiedN.Neighbors = append(copiedN.Neighbors, clone)
				continue
			}

			copiedN.Neighbors = append(copiedN.Neighbors, dfs(neigh))
		}

		return copiedN
	}

	res := dfs(node)
	return res
}

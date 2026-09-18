package generalalgonotes

type UnionFind struct {
	root []int
}

func newUnionFind(size int) *UnionFind {
	//sourcing behavior
	newRoot := make([]int, size)
	for i := range size {
		newRoot[i] = i
	}

	return &UnionFind{
		root: newRoot,
	}
}

func (uf *UnionFind) find(x int) int {
	return uf.root[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX := uf.find(x)
	rootY := uf.find(y)

	if rootX != rootY {
		for i := range uf.root {
			if uf.root[i] == rootY {
				uf.root[i] = rootX
			}
		}
	}
}

func (uf *UnionFind) connected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

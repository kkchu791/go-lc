package numberofprovinces

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}

	rank := make([]int, size)

	return &UnionFind{
		parent: parent,
		rank:   rank,
	}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	uf.link(uf.find(x), uf.find(y))
}

func (uf *UnionFind) link(x, y int) {
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
	} else {
		uf.parent[x] = y

		if uf.rank[x] == uf.rank[y] {
			uf.rank[y]++
		}
	}
}

func findCircleNumUnionFind(isConnected [][]int) int {
	size := len(isConnected)
	uf := NewUnionFind(size)
	numOfComponents := size

	for i := range size {
		for j := range size {
			if isConnected[i][j] == 1 && uf.find(i) != uf.find(j) {
				numOfComponents--
				uf.union(i, j)
			}
		}
	}

	return numOfComponents
}

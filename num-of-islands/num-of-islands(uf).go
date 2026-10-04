package numofislands

type UnionFind struct {
	Parents []int
	Ranks   []int
}

func newUF(size int) *UnionFind {
	parents := make([]int, 0)
	for n := range size {
		parents = append(parents, n)
	}

	ranks := make([]int, size)

	return &UnionFind{
		Parents: parents,
		Ranks:   ranks,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parents[x] != x {
		uf.Parents[x] = uf.Find(uf.Parents[x])
	}
	return uf.Parents[x]
}

func (uf *UnionFind) Union(x, y int) {
	xRoot := uf.Find(x)
	yRoot := uf.Find(y)

	if uf.Ranks[xRoot] > uf.Ranks[yRoot] {
		uf.Parents[yRoot] = xRoot
	} else {
		uf.Parents[xRoot] = yRoot
		if uf.Ranks[xRoot] == uf.Ranks[yRoot] {
			uf.Ranks[yRoot]++
		}
	}
}

func numIslandsUF(grid [][]byte) int {
	// first thing is lets create the union find
	size := len(grid) * len(grid[0])
	uf := newUF(size)
	cols := len(grid[0])

	// second thing is loop through the matrix
	// initialize  count

	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			cc := grid[r][c]

			if cc == '1' {
				count++
			}
		}
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				isOut := r < 0 || r >= len(grid) || c+1 < 0 || c+1 >= len(grid[0])

				if !isOut && grid[r][c+1] == '1' {
					u := r*cols + c
					v := r*cols + c + 1
					if uf.Find(u) != uf.Find(v) {
						uf.Union(u, v)
						count--
					}
				}

				// try down
				isOut = r+1 < 0 || r+1 >= len(grid) || c < 0 || c >= len(grid[0])

				if !isOut && grid[r+1][c] == '1' {
					u := r*cols + c
					v := (r+1)*cols + c

					if uf.Find(u) != uf.Find(v) {
						uf.Union(u, v)
						count--
					}
				}
			}

		}
	}

	return count
}

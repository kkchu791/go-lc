type UF struct {
    Rep []int
    Rank []int
}

func newUF(size int) *UF {
    rep := make([]int, 0)
    for n := range size {
        rep = append(rep, n)
    }

    rank := make([]int, size)

    return &UF{
        Rep: rep,
        Rank: rank,
    }
}

func (uf *UF) Find(x int) int {
    if x != uf.Rep[x] {
        uf.Rep[x] = uf.Find(uf.Rep[x])
    }
    return uf.Rep[x]
}

func (uf *UF) Union(x, y int) {
    x = uf.Find(x)
    y = uf.Find(y)

    if uf.Rank[x] > uf.Rank[y] {
        uf.Rep[y] = x
    } else {
        uf.Rep[x] = y

        if uf.Rank[x] == uf.Rank[y] {
            uf.Rank[y]++
        }
    }
}

func findRedundantConnection(edges [][]int) []int {
    size := len(edges)
    uf := newUF(size)

    for _, edge := range edges {
        u := edge[0] - 1
        v := edge[1] - 1

        if uf.Find(uf.Rep[u]) == uf.Find(uf.Rep[v]) {
            return edge
        } else {
            uf.Union(u, v)
        }
    }

    return nil 
}type UF struct {
    Rep []int
    Rank []int
}

func newUF(size int) *UF {
    rep := make([]int, 0)
    for n := range size {
        rep = append(rep, n)
    }

    rank := make([]int, size)

    return &UF{
        Rep: rep,
        Rank: rank,
    }
}

func (uf *UF) Find(x int) int {
    if x != uf.Rep[x] {
        uf.Rep[x] = uf.Find(uf.Rep[x])
    }
    return uf.Rep[x]
}

func (uf *UF) Union(x, y int) {
    x = uf.Find(x)
    y = uf.Find(y)

    if uf.Rank[x] > uf.Rank[y] {
        uf.Rep[y] = x
    } else {
        uf.Rep[x] = y

        if uf.Rank[x] == uf.Rank[y] {
            uf.Rank[y]++
        }
    }
}

func findRedundantConnection(edges [][]int) []int {
    size := len(edges)
    uf := newUF(size)

    for _, edge := range edges {
        u := edge[0] - 1
        v := edge[1] - 1

        if uf.Find(u) == uf.Find(v) {
            return edge
        }

    	uf.Union(u, v)
    }

    return nil 
}
package accountsmerge

import "sort"

type UF struct {
	repList  []int
	rankList []int
}

func newUF(size int) *UF {
	repList := make([]int, 0)
	for n := range size {
		repList = append(repList, n)
	}

	rankList := make([]int, size)

	return &UF{
		repList:  repList,
		rankList: rankList,
	}
}

func (u *UF) Find(x int) int {
	if x != u.repList[x] {
		u.repList[x] = u.Find(u.repList[x])
	}

	return u.repList[x]
}

func (u *UF) Union(x, y int) {
	u.Link(u.Find(x), u.Find(y))
}

func (u *UF) Link(x, y int) {
	if u.rankList[x] > u.rankList[y] {
		u.repList[y] = x
	} else {
		u.repList[x] = y
		if u.rankList[x] == u.rankList[y] {
			u.rankList[y]++
		}
	}
}

func accountsMergeUF(accounts [][]string) [][]string {
	size := len(accounts)
	uf := newUF(size)
	eg := make(map[string]int)

	for groupNum, acc := range accounts {
		for _, email := range acc[1:] {
			if _, exists := eg[email]; !exists {
				eg[email] = groupNum
			} else {
				uf.Union(groupNum, eg[email])
			}
		}
	}

	c := make(map[int][]string, 0)

	for email, groupNum := range eg {
		rep := uf.Find(groupNum)
		c[rep] = append(c[rep], email)
	}

	res := make([][]string, 0, len(c))
	for rep, emails := range c {
		sort.Strings(emails)
		name := accounts[rep][0]
		merged := []string{name}
		merged = append(merged, emails...)
		res = append(res, merged)
	}

	return res
}

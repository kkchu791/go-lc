package accountsmerge

import (
	"maps"
	"slices"
)

func accountsMerge(accounts [][]string) [][]string {
	//1st part create adj list

	// they all have to point to each other
	// i think it'll be anested loop

	al := make(map[string][]string, 0)
	res := make([][]string, 0)

	for _, account := range accounts {
		account = account[1:]

		for i := 0; i < len(account); i++ {
			uAccount := account[i]

			for j := 0; j < len(account); j++ {
				vAccount := account[j]

				if uAccount == vAccount {
					continue
				}

				al[uAccount] = append(al[uAccount], vAccount)
			}
		}
	}

	seen := make(map[string]bool, 0)
	visited := make(map[string]bool, 0)

	var getLinkedEmails func(string)

	getLinkedEmails = func(email string) {
		if seen[email] {
			return
		}

		visited[email] = true
		seen[email] = true

		for _, neigh := range al[email] {
			getLinkedEmails(neigh)
		}
	}

	for _, account := range accounts {

		if !seen[account[1]] {
			nameSlice := []string{account[0]}
			getLinkedEmails(account[1])

			emails := slices.Collect(maps.Keys(visited))
			slices.Sort(emails)

			newAccount := append(nameSlice, emails...)
			res = append(res, newAccount)
			// reset visited here:
			visited = make(map[string]bool, 0)
		}

	}

	return res
}

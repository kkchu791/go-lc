package keysandrooms

func canVisitAllRooms(rooms [][]int) bool {
	al := make(map[int][]int, 0)
	for idx, roomKeys := range rooms {
		al[idx] = roomKeys
	}

	var dfs func(int)

	visiting := make(map[int]bool, 0)
	visited := make(map[int]bool, 0)

	dfs = func(room int) {
		if visiting[room] {
			return
		}

		if visited[room] {
			return
		}

		visiting[room] = true

		for _, neigh := range al[room] {
			dfs(neigh)
		}

		delete(visiting, room)
		visited[room] = true
	}

	dfs(0)

	if len(visited) == len(rooms) {
		return true
	} else {
		return false
	}
}

func offLineMin(S[]int, n int) [] int {
    block := make ([]int, n + 1)

    count := 1
    for _, op := range S {
        if op == 0 {
            count++
        } else {
            block[op] = count
        }
    }
    m := count - 1

    nodes := make([]*Node, m + 2)

    for b := 1; b < m + 1; b++ {
        nodes[b] = &Node(val: b)
        makeset(nodes[b])
    }

    extracted := make([]int, m + 1)
    for i := 1; i <= n; i++ {
        j := find(nodes[block[i]]).val
        if j!= m + 1 {
            extracted[j] = i
            label = find(nodes[j+1]).val
            root := union(nodes[j], nodes[j+1])
            root.val = 1
        }
    }

    return extracted[1:]
}


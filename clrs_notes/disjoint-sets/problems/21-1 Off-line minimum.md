
### 21-1 Off-line minimum
input is S
m = number of Es
n = number of keys

j = 1
block = []
for each op in S {
    if op == E
        j++
    else
        block[op] = j
}

// make a set
for el in block {
    uf.Make-set(el)
}

make-set(x *Node)
    x.parent = x
    x.rank = 0

for b = 1; b <= m+1; b++ {
    nodes[b] = &Node{val: b}
    make-set(nodes[b])
}

extracted = []

for i = 1; i <= n; i++ {
    //find its current block
    j = label[Find(block[i])]
    if j != m+1 {
        extracted[j] = i
        l = label[Find(j+1)]  
        r = Union(j, j+1)
        label[r] = l
    }
}
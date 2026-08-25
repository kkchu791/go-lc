### Go Zero Value Semantics

```go

al := make(map[int][]int)

//edges are [[0, 1]]
for _, edge := range edges {
    al[edge[0]] = append(al[edge[0]], edge[1])
}

// output is:

// al := {
//     0: [1]
// }

```

this works because al[edge[0]] doesn't exist, and the zero value of 
slices are `[]`, which you can append to.

Make is mandatory for this to work because it allocates a valid non-nil map you can write to.

Declaring a variable map `var al map[int][]int` gives you a nil.
It has type map but its just nil.

So above adj list example works because we have a writeable hashmap from make.

The append works because `al[edge[0]]` returns nil, and `append` treats that as a regular slice with length 0.

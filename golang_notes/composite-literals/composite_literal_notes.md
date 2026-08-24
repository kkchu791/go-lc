### Composite Literal Notes

Composite literals are a way to define data structures like maps structs slices, and pass in the values like:

Values are usually passed in with {} 

Type + {Content}

```
[]int{1} //[1]

type Orange struct {
    taste string
}

o := Orange{
    taste: "sour"
}

//map
al := map[int][]int{
    0: []int{1},
    1: []int{2},
}
```

The option over composite literals is using `make`, which is more for initializing the data structure but not passing in contents just yet. You can add in the length, capacity if you know how much it'll take, but you won't add in the data yet.

// make examples

//slices
make (type, length(how many elements starts with (with zero fillers)), capacity)

make([]int, 0, 10) // slice of ints that can hold 10 elements, and is empty.

make([]int, 1, 5) // slice of ints that can hold 5 elements at most, and has 1 element of 0 inside of it (it fills always with zero fillers.)

//maps
make(map[int][]int, 10) map that can hold 10 elements

//chan
make(chan int, 10) channel that can send 10 elements, 11th element will be blocking.

Note: if you wanted it to fill with something else, you'd have to use composite literals.

Note: Make doesn't work for structs or arrays


### First Notes Iteration 1 (youtube video)

parent[i] = i
rank[i] = 1

make-union(1, 2)
parent 1 = find(1)
parent 2 = find(2)

```
find(x):
	while (x != parent[x])
		parent[x] = parent[parent[n]]
			x = parent[x]

	return x
```

Disjoint set data structure - stores a collection of non overlapping set
						- stores a partition of a set into a disjoint subsets

operations:
	1) adding new sets
	2) merging sets (replacing them by their union)
	3) finding a representative member of a set

Find - determine which subset a particular element is in. This can be used to determine if 2 elements are in the same subset.

Union - Join 2 subsets into a single subset.

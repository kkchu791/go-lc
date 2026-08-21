# Depth First Search Notes

Q:Why is predecessor v.Pi? It's the predecessor attribute.

Predecessor subgraph
Gpi = (V, En)

Predecessor subgraph  of DFS forms a **depth first forest** comprising several **depth first trees**.

When you dfs through a graph, you're also taking note of the prodecessor, almost like leaving bread crumbs. You can trace back your path back to the root. 

The edge of Epi are tree edges.

Logical clock goes up 1 whenever DFS discovers a new vertex or it finishes (backtracks). 

DFS timestamps its vertexes (Timestamps are done with a logical clock, not physical clock)
1. v.d (discovery) records when v is first discovered
2. v.f (finish) records when search finishes examining v's adjList and blackens v)

Vertex is white

	  |
	  v

1st timestamp: v.d records
	
	|
    v

Vertex is gray

2nd timestamp: v.f records

	|
	v

Vertex is black


```
DFS (G)

	for each vertex u (set of) G.V
		u.color = WHITE
		u.Pi = NIL
		
	time = 0
	for each vertex u (set of) G,V
		if u.color == WHITE
			DFS-visit(G.u)
```

```
DFS-visit(G, u)
	time = time + 1
	u.d = time
	u.color = GRAY

	for each v (set of) G.Adj[u]
		if v.color == WHITE
			v.Pi = u
			DFS-visit(G, v)
			
	u.color = BLACK
	time = time + 1
	u.f = time

```

Running time of DFS is O(V + E)

**Properties of DFS**

1) It yields valuable info about the structure of a graph
2) subgraph Gpi forms a forest of trees

Parenthesis structure - discovery and finishing times

Structure of depth first trees exactly mirrors the structure of recursive calls of DFS-visit



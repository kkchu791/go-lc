# Breadth First Search Notes

Some good ones to try:
- Prim's minimum spanning tree algorithm
- Dijkstra's sing source shortest path algorithm

Formula for a Graph
G = (V, E)

**Source** - starting point

BFS works on both directed and undirected graphs.

BFS "expands the frontier" between discovered and undisdcovered vertices uniformly across the breadth of the frontier.

BFS colors each vertex white, gray, or black
Discovered makes it non white
Gray is discovered but not visited? They are in Queue
Black is visited

BFS constructs a breadth first tree
   - root is the source vertex S.

```
BFS(G, s)
  for each vertex u (subset of) G.V - {s}
	  u.color = WHITE
	  u.d = INFINITY
	  u.Pi = NIL
	  
  s.color = GRAY
  s.d = 0
  s.PI = NIL
  Q = empty
  Enqueue(Q, s)
  while Q is not empty
	  u = Dequeue(Q)
	  for each v (subet of) G.Adj[u]
		  if v.color == WHITE
			v.color = GRAY
			v.d = u.d + 1
			v.Pi = u
			EnQueue(Q,v)
			
	  u.color = BLACK
```

Note: Gray vertices are not had their adj list examined
Run time: O(V+E)

**Lemma 22.1 Small subt theorem** - shortest path will always be less than shortest path + 1 (duh!)
  Shortest Path(s, v) <= Shortest path(s, u) + 1

**Corollary**
  Proves that BFS finds the shortest path distance


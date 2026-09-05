### Topological Sort Iter 2

It's a linear ordering of all its vertices, such that if G contains an edge (u, v), the un appears before v in the ordering.
![[Screenshot 2026-08-31 at 4.00.24 PM.png|700]]


![[Screenshot 2026-08-31 at 4.00.36 PM.png|700]]

Topo Sort -  u appears before v
		Ordering of its vertices along a horizontal line, edges go left to right

```
Topoloical-Sort(G)
1 call DFS(G) to compute finishing time v.f for each vertex v
2 as each vertex is finished, insert it onto the front of a linked list
3 return the linked list of vertices
```

Topo Sort takes O( V + E) Run time

Lemma 22.11 - Directed graph is acyclic if it has no back edges

White path theorem - In DFS graph, v is a descendant of u if at time time of discovery of u, there is a path from u  to v consisting entirely of white vertices.

Seems like its just proving if you have a back edge, you have a cycle. So u goes to v, and v goes to some ancestor of u

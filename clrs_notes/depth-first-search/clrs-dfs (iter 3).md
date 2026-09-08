### CLRS-DFS

backtracks - bubbles, propogates

In BFS, the predecessor subgraph forms a tree (every node has exactly one parent)
In DFS the predecessor forms several trees

v.d - gray
v.f - black

BFS - serves to find shortest path
DFS - usually serves as a a subroutine in another algorithm

For every vertex u, u.d < u.f

O(E) -> Total cost of exploring the neighbors
O(V) -> Iterating through the nodes

DFS is O(E+V)

**2 Noteable Properties for DFS**
1. Predecessor subgraph Gpi forms a forest of trees.

	Struct of Depth First Trees mirrors recursive calls of DFS Visit.

- v is a descendant of u if exploring while u is gray(visiting)
- u equals v's predecessor **if and only** if DFS-Visit(G, v) called during a search of u's adj list

![[Screenshot 2026-09-07 at 5.59.26 PM.png|582]]

	DFS visit is called twice in the algo, one to start the tree search and one to explore the neighbors. When exploring neighbor v's predecessors = u. If your starting to tree, the root has no predecessor

2. v.d and v.f have parenthesis structure

	Theorem 22.7 (Parenthesis Theorem)
![[Screenshot 2026-09-07 at 6.02.30 PM.png]]
	If any DFS of a directed or undirected graph, for u and v, exactly one of the following 3 conditions hold:

	1. if interval [u.d, uf] and interval [v.d, v.f] are entirely disjoint (no overlaps), neither u nor v is a descendant of the other
	2. if interval [u.d, u.f] is contained entirely with [v.d, v.f] , u is a descendant of v in a depth first tree.
	3. if interval [v.d, v.f] is contained entirely within [u.d, u.f], v is a descendant of u in a depth first tree.


	
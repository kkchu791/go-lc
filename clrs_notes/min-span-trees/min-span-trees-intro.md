### Minimum Spanning Trees

w (u, v) - weight - the cost (amount of wire needed) to connect u and v

Goal: find an acyclic subset T(tree) set of E(edges) that connects all of the vertices

total weight is minimized 

Since T is acyclic, connects all of the vertices, it forms a tree.
We call it a **spanning tree**, since it "spans" the graph G

How to figure min span tree? Greedy Algo?

Prim's Algo - resembles Dijkstra's shortest path algo
Kruskal Algo - similar to connected components algo

Greedy - making the choice that is the best at the moment

![[Screenshot 2026-08-29 at 4.21.11 PM.png]]For min spanning tree probems, we can prove that certain greedy strategies do yeild a spanning tree min weight.

23.1: Growing a min span tree
G= (V, E)

weight func W: E -> R

Loop invariant:

Prior to each iteration, A is a subset of some min spanning trees.

A (set union) {(u, v)} 
A - generic method manages a set of edges A

Safe Edge - an edge we can add to A because we can add it safely to A.

```
Generic MST(G, w)
	A = not empty
	while A does not form a spanning tree
		find an edge(u, v) that is safe for A
			A = A (set union) {(u, v)}
			
	return A
```
 
 Loop invariant:
 Initialization - set A satisfies the loop invariant
 Maintenance - loop maintains the invariant by adding only safe edges
 Termination - All edges added to A are in a minimum spanning tree and so set A must be a min spanning tree min spanning tree == set of min spanning trees

A cut(S, V-S) of an undirected graph G= (V, E) is a partition of V.

An edges **crosses** the cut if one endpoint is in S and the other in V-S.

A cut **respects** a set A of edges if no edge in A cross a cut.

An edge is a light edge, crossing a cut if its weight is the min of any edge crossing the cut.

One rule for recognizing safe edges is given:

Black vertices are in the set S
White vertices are in V-S
Edges crossing the cut are those connecting white vertices w/ black vertices

Edge DC is the unique light edge crossing the cut
A subset A of the dges is shaded

Note: the cut (s, v-S) respects A since edge of A crosses the cut


### Strongly Connected Components

Its a classic application of DFS.

Decomposing a directed graph into its strongly connected components.

Use 2 DFFs

G = (V, E) is a maximal set of vertices C (set of) V

Strongly connected components of a directed graph G = (V, E) is a maximal set of vertices C (set of V)

C = representation of strongly connected components


![[Screenshot 2026-08-30 at 4.01.19 PM.png]]
```
Strongly-Connected-Components(G)
1. call DFS(G) to compute finishing time u.f for each vertex u
2. compute G^T (graph transpose)
3.call DFS(G^T), but in the main loop of DFS, consider the vertices in order of decreasing v.f
3. output the vertices of each tree in the depth first forst as a seperate strongly connected component.
```

Component Graph
	G^scc = (V^scc, E^scc)

Lemma 22.13 subtheorem

v` ~ v ~ u

The point of SCC is to find "maximal" groups of nodes that are all mutually reachable from each other.

SCC algo are for grouping connected nodes into one vertex so you can treat the Graph as a DAG.


**Some use cases**

Compilers - detecting circular dependencies. Compile the SCC dependencies as one unit.
Social Media Graphs - used for seeing which users are in tightly linked communities


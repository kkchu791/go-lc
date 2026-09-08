### Classification of Edges

Search can be used to classify the edges.
Type of edge can provide important info about a graph
Direct graph is acyclic **if and only if (iff)** a DF search yields no back edges (lemma 22.11)

Tree edges - edges in the depth first forest
			Edge (u, v) (neighbor) is a tree edge if v was first discovered by exploring edge (u,v)
Back edges - edges (u, v) connecting a vertex u to an ancestor v in a depth first tree. Self loops

Forward Edges - like a shortcut. Non-tree edge

![[Screenshot 2026-09-08 at 2.19.06 PM.png|462]]

Forward Edges represents actual existing connections in a graph

Cross Edges - go between vertices in the same depth frist tree as along as one vertex is not an ancestor of the other or go between different vertices in different depth first trees.

![[Screenshot 2026-09-08 at 2.19.13 PM.png|443]]
Undirected graph 
- Forward and cross edges never occur in undirected graphs

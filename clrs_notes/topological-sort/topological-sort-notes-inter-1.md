# Toplogical Sort

u appears before v
Order of its vertices along a horizontal line, edges go left to right

A -> B -> C

Professor Bumstead toplogically sorts his clothing when getting dressed

Vertices arranged from left to right in order of decreasing finishing time.

```
Toplogical-Sort(G) (simple algo for sorting a dag)
1. call DFS(G) to compute finishing times of v.f for each vertex v
2. As each vertex is finished, insert it onto the front of a linked list
3. Return the linked list of vertices

```

DFS take O(V+E) and O(1) to insert each of the |V| onto the front of linked lists.

**Lemma 22.1**

Graph is acyclic if dfs yields no back edges (cycles)

Proof -> v connects back to u

"At the time v.d, the vertices of c form a path of white vertices from v to u"
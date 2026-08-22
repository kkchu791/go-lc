# Intuition
i was thinking canFinish as the dfs
I was thinking if you can get through the whole dfs without hitting a cycle, that means can return the dfs path
if you hit a cycle, you return empty array.
Also thought about, what if you added to the list as you went down.
If you hit a cycle, you start deleting from the list as you bubble up
I thought that give me the same answer, rather than using a global list
which felt like could get messy.
I didn't think to populate the list as you bubble up, so you're like safe when you're bubblying up because you've reach the end of the dfs, and you know you don't have a cycle.
I didn't think about that.

# Approach
global res list and bubbling up and populating the res list once you know you don't have a cycle.
Use visiting to detect cycle.
Use visited to memo optimize.
Adjacency List to search the next courses to take.


# Complexity
- Time complexity:
O(V+E) -> O(N) // searching every vertex just once
- Space complexity:
O(V+E -> O(N) // you add an adj list with all edges and vertex.
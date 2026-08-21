### Intuition
Bubble up 1 from each node
if node nil return 0
Compare left and right to get max and send that up + 1 //bestLine + 1 more
return the end results from the dfs

### Approach
I could solve so I looked at a popular answer.
When I saw the use of a global variable to track best diameter.
I ran through it on whitebard and saw it could work.
My intuition was correct to get the best line (best Left or best Right)
to keep sending up to see if there are better,
I never thought of calculating best diameter at each node as a global vairable
but it works. Its tricky.

### Complexity
Time complexity:
O(n), we have through every vertex on the graph.

Space complexity:
O(h) -h being heigh of the tree. Itsfor the recursive stack.
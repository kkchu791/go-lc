### Intuition
Go right and down with dfs
Update a global counter when you reach target
Handle for out of bounds
Memo Seen so you don't have to go through all previous cells again

### Approach
Kept intuition approach, but didn't consider that a cell could have multiple path (i was counting just 1). So I changed it propagate up the number of paths when you hit the target.

So each cell knows how many paths are possible.

Then I memoized this, so if you'd hit the previous cell, you'd know how many paths are possible.

### Complexity
Time complexity:
O(m * n) // 3 *3 you'd go through 9 cells

### Space complexity:
O(m + n) // 3 *3 you'd save 9 cells in your memo


### Intuition
adj list
loop through numCoursees
detect if there's a cycle
save results to shorten dfs
use visiting for cycle detection

### Approach
same as my intuition, but the only think i didn't really understand is how to create the adjlist

it should be: course you need to take -> [to take these courses]

i thought it was: for this course -> [you need to first take these]

ugh.

### Complexity
Time complexity:
O(V+E)
Space complexity:
O(V+E)

### Intuition
Union Find approach:
You have to give it size to determine the parents and ranks list.
The methods are union and find (path compression)
Its preety cool
The only thing is the nest loop of the graph is something I wouldn't really get on my own. You could just loop through the whole matrix,
Trying to find bridges to existing components
The two conditions for when he have numC--
were:
currCell had to be a 1 and the row column element has to have different parents (if same parent, it means they are in the same component)

This means we've found a road to an existing component. We can then union the two and decrease the our possible component count.

### Complexity
Time complexity:
we loop through size so thats O(size)
we loop through size again so that O()
matrix was O(r + c), it would half of this
union has find in it so that is possibly O(parent)
find is the same

i think its this since we do 2 loops O(n²)

Space complexity:
Space is linear.
# Intuition /Approach

I knew topo sort going in but still didn't understand how to use it for this question.
Seems like knowing the rules helped, you know, the rules are defined by first word compared to second word
specifically each character in these words
If there's a mismatch, tt becomes an edge.

    The premise here is the problem, the "Aliens" are giving us a code and its presumably sorted in their understanding of sorting. So 

    For instance "z", "o", in our language z is larger than o, but in their language z is smaller than o
    because they ordered it that way, z coming before o means its smaller.

    Same with hrn and hrf, in the alien language n comes before f.

    We add it to a graph to show that n has to come before f, same with z coming before o

    We create all the pre reqs for the letter that we know of and then we topo sort it

    Topo sort means we are sorting by finishing time, which means once we finish this letter, we know the other letters after can be finished next, they will be sorted in the correct finishing order.


and it keeps going, that is lexicographical ordering
You topo sort the graph to get the ordering. (I'm not sure why this works)

The other part is the invalid rules
Invalid if len(A) is greater than len(B)
so we handle that, if theres no mismatch and lenght A is greater than length B, we return ""



# Time/ Space Complexity

lets talk about time:  
we loop through words
    we also loop through every character in those words 

O(N * M) which is O(C)

we iterate through our set of letters
iterate through their edges but we memoized so we only lopp through edges once

i think O(V+E), memo stopped it from being exponential


We also reverse the slice, O(n)

overall i think its O(C) + O(V+E) +  O(n) = so O(C) simplified

For space:
we have a set from the letters (O(n))
we have a al from the letters (O(n))
we have a visited set (o(V))
we have a visiting set (o(E))
we have a visitedSlice set (o(V))

We only store up to 26 letters so its a constant time of O(1)
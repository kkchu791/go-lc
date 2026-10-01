Intuition
Hmm so Union Find:
Seems like the first thing you need to do is create an email group, key is the email and it points to the groupNum. If the email is already in the list we union the current groupNum of the email with the first group num the one we put in the groupNum map.

So we do some unions and we have this emailGroup map, which has all the emails in it, and they are pointing to a groupNum that may or may not be the representative of the group, but we'll get to the true representative in the next loop.

We loop through the email group. We call a find-set on the groupNum, which will return the true representative of the component, and we add it another map, lets call it components.
By the end we'll have a map of emails that are all grouped together if they share the same rep.
By the find method, we also do path compression, which I believe saves on time complexity.(Need to check)

Afterwords we just have to append the name to the emails list and return it in a res list.

Complexity
N = accounts, K = max emails per account

Time complexity:
Time: O(NK log(NK)). Unions and finds are O(NK · α(N)), so the sorting dominates.

Space complexity:
Space: O(NK), not O(N). The maps hold every email. Only the DSU arrays are O(N).
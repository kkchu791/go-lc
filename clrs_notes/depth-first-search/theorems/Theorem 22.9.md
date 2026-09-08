Theorem 22.9 (White-path theorem)


Vertex v is a descendant of vertex u if and only if when you discover u, there is a path from u to v consisting entirely of white vertices

**Proof =>** (P imples Q)

If v = u, path from u to v contains just vertex u, which is still white when we set the value of u.d.

Suppose v is descendant of u.
By Corollary 22.8  = u.d < v.d, v is white at the time of u.d
Since v can be any descendant of u, all vertices on the unique simple path from u to v in the depth first forest are white at time u.d.![[Screenshot 2026-09-07 at 6.28.28 PM.png]]
**Proof <=** (Proof by Contradiction, Reductio ad absurdum)
w is v predecessor
u.d < v.d < w.f <= u.f

By **Parenthesis Theorem** (Theorem 22.7), interval [v.d, v.f] is entirely contained with [u.d, uf] which forces v to actually be a descendant of u.


Proof by contradiction
a) Assume the exact opposite of what you want to prove
b) Follow Rules and theorems until you reach something impossible
c) Conclude your initial opposite assumption must be false which means the original statement must be true.

Goal: Prove v must become a descendant of u

Setup: v does not become a descendant of u.

Chain of Logic:
 - v is not discovered
 - v is on a path of white vertices
 - v has to be discovered after u (u.d < v.d)

- v's predecesssor w does not become a descendant of u, so w finishes before u (w.f <= u.f)

- w is connected o v
- v is white when u is discovered
- v must be discovered before w finishes (u.d < w.f)

	u.d < v.d < w.f <=  u.f

The Contradiction: v was discovered during u's interval. By Parenthesis Theorem, any vertex discovered while u is active if a descendant of u.

Conclusion: Our assumption v is not a descendant led us to v must be a descendant.
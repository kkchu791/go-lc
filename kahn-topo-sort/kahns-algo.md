
### Kahns Algo

It seems like a bfs approach to doing topo sort

```
Kahn's Algo is topo sort on a DAG
1. Count indgrees for every node
2. Push all nodes with indegree of 0 to queue
3. Remove vertex from the queue, add to result list, reduce inD of all its adj vertices
4. Add the indegrees of 0 to the queue.
5. Process continues until the queue is empty
6. The result list is a topo sort.
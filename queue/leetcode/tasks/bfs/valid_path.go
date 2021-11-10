package main

func validPath(n int, edges [][]int, start int, end int) bool {
	if n == 1 {
		return true
	}

	graph := newGraph(edges)
	visited := make(map[int]struct{})

	var queue struct {
		Size     int
		Elements []int
	}

	neighbors, ok := graph[start]
	if ok {
		visited[start] = struct{}{}
	} else {
		visited[end] = struct{}{}
		neighbors = graph[end]
	}

	queue.Size = len(neighbors)
	queue.Elements = neighbors
	queueNum := -1

	for queue.Size != 0 {
		queueNum++
		queue.Size--

		currNode := queue.Elements[queueNum]
		if currNode == end {
			return true
		}

		visited[currNode] = struct{}{}
		for _, neighbor := range graph[currNode] {
			if _, ok := visited[neighbor]; ok {
				continue
			}

			queue.Size++
			queue.Elements = append(queue.Elements, neighbor)
		}
	}

	return false
}

func newGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int, len(edges))

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return graph
}

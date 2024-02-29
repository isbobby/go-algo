package topology

import (
	"sort"
)

func TopologicalSort(edges map[int][]int, nodes []int) []int {
	visitedMap := make(map[int]bool, len(nodes))
	topologicallySortedNodes := []int{}

	// DFS and append at the end
	var dfsAndAppend func(v int, neighbors []int)
	dfsAndAppend = func(v int, neighbors []int) {
		for i := range neighbors {
			visited, exists := visitedMap[neighbors[i]]

			// only do DFS if we have not visited the node
			if !visited || !exists {
				visitedMap[neighbors[i]] = true
				dfsAndAppend(neighbors[i], edges[neighbors[i]])
			}
		}

		topologicallySortedNodes = append(topologicallySortedNodes, v)
	}

	for i := range nodes {
		visited, exists := visitedMap[nodes[i]]

		// only do DFS if we have not visited the node
		if !visited || !exists {
			visitedMap[nodes[i]] = true
			dfsAndAppend(nodes[i], edges[nodes[i]])
		}
	}

	// Reverse the order at the end
	// As the leaf nodes will be the first one to be appended
	sort.Slice(topologicallySortedNodes, func(i, j int) bool {
		return i > j
	})

	return topologicallySortedNodes
}

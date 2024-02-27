package mst

import (
	"sort"
)

// problem taken from leetcode l1584

// given some points on a cartesian plane

// return the minimum cost to make all points connected

// distance between points is manahttan distance

func minCostConnectPoints(coordinates [][]int) int {
	points := []point{}
	for i := range coordinates {
		points = append(points, point{
			coordinates[i][0],
			coordinates[i][1],
		})
	}

	edeges := edegeList{}
	// calculate edeges, N sqaured
	var computeEdegesBackTrack func(currentPoint point, points []point)
	computeEdegesBackTrack = func(currentPoint point, points []point) {
		if len(points) == 0 {
			return
		}

		for i := 0; i < len(points); i++ {
			edeges = append(edeges, edge{a: currentPoint, b: points[i], w: manhattanDistance(currentPoint, points[i])})
		}

		computeEdegesBackTrack(points[0], points[1:])
	}
	computeEdegesBackTrack(points[0], points[1:])

	// sort edeges
	sort.Sort(edeges)

	parents := make(map[point]point, len(points))
	// build a forest of single node trees
	for i := range points {
		parents[points[i]] = points[i]
	}

	totalWeight := 0
	// select from sorted edges to combine trees
	for i := range edeges {
		edge := edeges[i]

		if !pointsSameRoot(edge.a, edge.b, parents) {
			unionPoints(edge.a, edge.b, parents)
			totalWeight += edge.w
		}
	}

	return totalWeight
}

func getPointsRoot(p point, parent map[point]point) point {
	if parent[p].equal(p) {
		return parent[p]
	}

	return getPointsRoot(parent[p], parent)
}

func pointsSameRoot(a, b point, parent map[point]point) bool {
	rootA := getPointsRoot(a, parent)
	rootB := getPointsRoot(b, parent)

	return rootA.equal(rootB)
}

func unionPoints(a, b point, parent map[point]point) {
	rootA := getPointsRoot(a, parent)
	rootB := getPointsRoot(b, parent)

	if rootA == rootB {
		return
	}

	parent[rootA] = parent[rootB]
}

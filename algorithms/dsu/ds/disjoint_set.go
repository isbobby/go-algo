package ds

import "fmt"

type DisjointSet struct {
	parents []int
	sizes   []int
}

func (s *DisjointSet) NaiveSameSet(a, b int) bool {
	return s.NaiveFindRoot(a) == s.NaiveFindRoot(b)
}

// return the root value of v
func (s *DisjointSet) NaiveFindRoot(v int) int {
	if s.parents[v] == v {
		return v
	}

	return s.NaiveFindRoot(s.parents[v])
}

func (s *DisjointSet) NaiveUnion(a, b int) {
	rootA := s.NaiveFindRoot(a)
	rootB := s.NaiveFindRoot(b)

	if rootA == rootB {
		return
	}

	s.parents[b] = rootA
}

func (s *DisjointSet) MakeSet(v int) {
	s.parents[v] = v
}

func (s *DisjointSet) Show() {
	fmt.Printf("Root :")
	for i := range s.parents {
		fmt.Printf("%v ", s.parents[i])
	}

	fmt.Printf("\nIndex:")
	for i := range s.parents {
		fmt.Printf("%v ", i)
	}
	fmt.Println("\n-----------------")
}

func InitializeDisjointSet(size int) *DisjointSet {
	return &DisjointSet{
		parents: make([]int, size),
		sizes:   make([]int, size),
	}
}

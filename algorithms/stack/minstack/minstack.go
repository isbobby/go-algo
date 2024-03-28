package minstack

type MinStack struct {
	minIntStack []int
	intStack    []int
}

func Constructor() MinStack {
	return MinStack{
		minIntStack: []int{},
		intStack:    []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.intStack = append(s.intStack, val)

	if len(s.minIntStack) == 0 {
		s.minIntStack = append(s.minIntStack, val)
	} else if val <= s.minIntStack[len(s.minIntStack)-1] {
		s.minIntStack = append(s.minIntStack, val)
	}
}

func (s *MinStack) Pop() {
	if s.intStack[len(s.intStack)-1] == s.minIntStack[len(s.minIntStack)-1] {
		s.minIntStack = s.minIntStack[:len(s.minIntStack)-1]
	}

	s.intStack = s.intStack[:len(s.intStack)-1]
}

func (s *MinStack) Top() int {
	return s.intStack[len(s.intStack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minIntStack[len(s.minIntStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

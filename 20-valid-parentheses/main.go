package validparentheses

type stack struct {
	values []int8
}

func (s *stack) Length() int {
	return len(s.values)
}

func (s *stack) Push(val int8) {
	s.values = append(s.values, val)
}

func (s *stack) Pop() int8 {

	if len(s.values) == 0 {
		return int8(-1)
	}

	result := s.values[len(s.values)-1]

	s.values = s.values[:len(s.values)-1]

	return result
}

func (s *stack) Peek() int8 {

	if len(s.values) == 0 {
		return int8(-1)
	}

	return s.values[len(s.values)-1]
}

func NewStack() *stack {
	return &stack{
		values: make([]int8, 0),
	}
}

func isValid(s string) bool {

	stk := NewStack()

	for _, char := range s {
		switch char {
		case '(':
			stk.Push(int8(1))
		case '[':
			stk.Push(int8(2))
		case '{':
			stk.Push(int8(3))
		case ')':
			if stk.Pop() != int8(1) {
				return false
			}
		case ']':
			if stk.Pop() != int8(2) {
				return false
			}
		case '}':
			if stk.Pop() != int8(3) {
				return false
			}
		}
	}

	return stk.Length() == 0
}

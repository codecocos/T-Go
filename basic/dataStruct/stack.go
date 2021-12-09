package dataStruct

type Stack struct {
	ll *LinkedList
}

func NewStack() *Stack {
	//새로운 링크드리스트를 만들어서 메모리주소를 넣어 초기화 하고 반환
	return &Stack{ll: &LinkedList{}}
}

func (s *Stack) Empty() bool {
	return s.ll.Empty()
}

func (s *Stack) Push(val int) {
	s.ll.AddNode(val)
}

func (s *Stack) Pop() int {
	back := s.ll.Back()
	s.ll.PopBack()
	return back
}
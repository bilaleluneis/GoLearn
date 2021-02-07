/*	@author: Bilal El Uneis
	@since: Feb 2021
	@email: bilaleluneis@gmail.com	*/

package internal

import types "GoLearn/src/common"

type Stack struct {
	head *node
}

func NewStack() Stack {
	return Stack{nil}
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Size() uint {
	if s.IsEmpty() {
		return 0
	}
	var counter uint = 1
	for currNode := s.head; currNode.next != nil; currNode = currNode.next {
		counter++
	}
	return counter
}

func (s *Stack) Push(item types.T) {

	if s.IsEmpty() {
		s.head = &node{nil, nil, item}
		return
	}

	currNode := s.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	currNode.next = &node{nil, nil, item}

}

func (s *Stack) Pop() (ok bool, value types.T) {

	if s.IsEmpty() {
		return ok, value
	}

	ok = true
	if s.head.next == nil {
		value = s.head.data
		s.head = nil
	} else {
		value = s.head.data
		top := s.head
		s.head = s.head.next
		top.next = nil
		top = nil
	}

	return ok, value
}

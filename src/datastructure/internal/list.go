/*	@author: Bilal El Uneis
	@since: Feb 2021
	@email: bilaleluneis@gmail.com	*/

package internal

import types "GoLearn/src/common"

//TODO: ensure that all elements are of same type use head.data.(type) == item.(type)
//check chapter 7 learning go .. type assertion vs typecasting
// castedTo, ok  = i.(int)

type List struct {
	head *node
}

func NewList() List {
	return List{nil}
}

func (l *List) IsEmpty() bool {
	return l.head == nil
}

func (l *List) Size() uint {
	if l.IsEmpty() {
		return 0
	}
	var counter uint = 1
	for currNode := l.head; currNode.next != nil; currNode = currNode.next {
		counter++
	}
	return counter
}

func (l *List) Push(item types.T) {

	if l.IsEmpty() {
		l.head = &node{nil, nil, item}
		return
	}

	currNode := l.head
	for currNode.next != nil {
		currNode = currNode.next
	}
	currNode.next = &node{nil, nil, item}

}

func (l *List) Pop() (bool, types.T) {

	var value interface{} = nil
	var success bool

	if l.IsEmpty() {
		return success, value
	}

	if l.head.next == nil {
		value, success = l.head.data, true
		l.head = nil
		return success, value
	}

	currHead := l.head.next
	currNext := currHead.next

	for currNext != nil {

		if currNext.next == nil {
			success, value = true, currNext.data
			currHead.next = nil
			break
		}

		currHead = currNext
		currNext = currHead.next
	}

	return success, value
}

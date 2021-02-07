/*	@author: Bilal El Uneis
	@since: Feb 2021
	@email: bilaleluneis@gmail.com	*/

package datastructure

import (
	types "GoLearn/src/common"
	"GoLearn/src/datastructure/internal"
)

type Collection interface {
	IsEmpty() bool
	Size() uint
	Push(value types.T)
	Pop() (success bool, value types.T)
}

func List() Collection {
	list := internal.NewList()
	return &list
}

func Stack() Collection {
	stack := internal.NewStack()
	return &stack
}

/*	@author: Bilal El Uneis
	@since: Feb 2021
	@email: bilaleluneis@gmail.com	*/

package test

import (
	"GoLearn/src/datastructure"
	"testing"
)

func TestCollection(t *testing.T) {
	c := datastructure.List()
	if c.Size() != 0 {
		t.Fail()
	}
}

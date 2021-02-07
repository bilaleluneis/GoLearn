/*	@author: Bilal El Uneis
	@since: Feb 2021
	@email: bilaleluneis@gmail.com	*/

package test

import (
	"GoLearn/src/immutability"
	"testing"
)

func TestOne(t *testing.T) {
	personAddr := immutability.NewAddress("111 w 6th")
	cantModify(personAddr)
	canModify(personAddr, "new address")
}

// helper functions for tests

func cantModify(address_ immutability.IMutAddress) string {
	return address_.Get()
}

func canModify(address_ immutability.Address, newAddr string) {
	address_.Update(newAddr)
}

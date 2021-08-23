/*	@author: Bilal El Uneis
	@since: August 2021
	@email: bilaleluneis@gmail.com	*/

package generics

import (
	"testing"
)

/* gnew : generic new() allows the use of T type to return heap allocated instance of T*/
func gnew[T interface{}]() *T {
	var instance T
	return &instance
}

func TestFunctions(t *testing.T) {
	// Test generic heap allocation
	if integer := gnew[int](); *integer != 0 {
		t.Fail()
	}
}

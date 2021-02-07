/*	@author: Bilal El Uneis
	@since: Jan 2021
	@email: bilaleluneis@gmail.com	*/

package internal

import types "GoLearn/src/common"

type node struct {
	prev *node
	next *node
	data types.T
}

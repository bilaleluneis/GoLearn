/*	@author: Bilal El Uneis
	@since: Feb 2021
	@email: bilaleluneis@gmail.com	*/

package immutability

type address struct {
	value string
}

func (addr *address) Update(newAddress string) {
	addr.value = newAddress
}

func (addr *address) Get() string {
	return addr.value
}

type IMutAddress interface {
	Get() string
}

type Address interface {
	IMutAddress
	Update(newAddress string)
}

func NewAddress(address_ string) Address {
	return &address{value: address_}
}

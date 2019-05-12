package intarray

type Iterator interface {
	Value() int
	Next() bool
}

type IntIterator struct {
	data    []int
	current int
}

func (it *IntIterator) Value() int {
	return it.data[it.current]
}

func (it *IntIterator) Next() bool {
	it.current++
	if it.current >= len(it.data) {
		return false
	}

	return true
}

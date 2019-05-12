package intarray

import (
	"fmt"
)

type IntArray struct {
	v []int
}

func New(i []int) *IntArray {
	a := IntArray{v: i}
	return &a
}

func (a *IntArray) Get(index int) int {
	return a.v[index]
}

func (a *IntArray) Set(index, value int) {
	a.v[index] = value
}

func (a *IntArray) Push(i int) {
	a.v = append(a.v, i)
}

func (a *IntArray) Min() int {
	return Min(a.v)
}

func (a *IntArray) Max() int {
	return Max(a.v)
}

func (a *IntArray) Val() []int {
	return a.v
}

func (a *IntArray) String() string {
	return fmt.Sprintf("%v", a.v)
}

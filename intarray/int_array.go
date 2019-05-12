package intarray

import (
	"fmt"
)

type IntArray struct {
	array []int
}

func New(i []int) *IntArray {
	a := IntArray{array: i}
	return &a
}

func (a *IntArray) Get(index int) int {
	return a.array[index]
}

func (a *IntArray) Set(index, value int) {
	a.array[index] = value
}

func (a *IntArray) Push(i int) {
	a.array = append(a.array, i)
}

func (a *IntArray) Min() int {
	return Min(a.array)
}

func (a *IntArray) Max() int {
	return Max(a.array)
}

func (a *IntArray) String() string {
	return fmt.Sprintf("%v", a.array)
}

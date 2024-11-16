package ordered_slice

import (
	"fmt"
	"sort"
)

type OrderedSlice[T comparable] struct {
	storage []T
	length  int
}

func (orderedSlice *OrderedSlice[T]) Add(value T) {
	index := orderedSlice.binarySearch(value)
	if index == len(orderedSlice.storage) { // no inserted value in slice, so we just append it to the end of storage
		orderedSlice.storage = append(orderedSlice.storage, value)
	} else if index == 0 { // need to create new storage, where first elem is inserted value, and others -> old storage
		temp := make([]T, orderedSlice.Len()+1)
		temp[0] = value
		copy(temp[1:], orderedSlice.storage)
		orderedSlice.storage = temp
	} else {
		// Comparing elem by searched index to inserted value. If inserted value is greater than value by index,
		// we need to increase index to copy left part of storage correctly
		// (including value by index before increasing).
		switch v := any(value).(type) {
		case int:
			if any(orderedSlice.storage[index]).(int) < v {
				index++
			}
		case string:
			if any(orderedSlice.storage[index]).(string) < v {
				index++
			}
		case float64:
			if any(orderedSlice.storage[index]).(float64) < v {
				index++
			}
		case float32:
			if any(orderedSlice.storage[index]).(float32) < v {
				index++
			}
		}

		temp := make([]T, orderedSlice.Len())
		copy(temp, orderedSlice.storage)
		orderedSlice.storage = append(orderedSlice.storage[:index], value)
		orderedSlice.storage = append(orderedSlice.storage, temp[index:]...)
	}

	orderedSlice.length++
}

func (orderedSlice *OrderedSlice[T]) Remove(value T) bool {
	index := orderedSlice.binarySearch(value)
	if index == len(orderedSlice.storage) || orderedSlice.storage[index] != value {
		return false
	}

	orderedSlice.storage = append(orderedSlice.storage[:index], orderedSlice.storage[index+1:]...)
	orderedSlice.length--
	return true
}

func (orderedSlice *OrderedSlice[T]) Contains(value T) bool {
	index := orderedSlice.binarySearch(value)
	if index == len(orderedSlice.storage) {
		return false
	}

	return orderedSlice.storage[index] == value
}

func (orderedSlice *OrderedSlice[T]) Index(value T) int {
	index := orderedSlice.binarySearch(value)
	if index == len(orderedSlice.storage) {
		return -1
	}

	for index-1 >= 0 && orderedSlice.storage[index-1] == value {
		index--
	}

	return index
}

func (orderedSlice *OrderedSlice[T]) At(index int) T {
	return orderedSlice.storage[index]
}

func (orderedSlice *OrderedSlice[T]) Len() int {
	return orderedSlice.length
}

func (orderedSlice *OrderedSlice[T]) Clear() {
	orderedSlice.storage = []T{}
	orderedSlice.length = 0
}

func (orderedSlice *OrderedSlice[T]) binarySearch(value T) int {
	left, right := 0, len(orderedSlice.storage)
	for left < right {
		mid := int(right+left) / 2
		if orderedSlice.storage[mid] == value {
			return mid
		}

		switch v := any(value).(type) {
		case int:
			if any(orderedSlice.storage[mid]).(int) < v {
				left = mid + 1
			} else {
				right = mid - 1
			}
		case string:
			if any(orderedSlice.storage[mid]).(string) < v {
				left = mid + 1
			} else {
				right = mid - 1
			}
		case float64:
			if any(orderedSlice.storage[mid]).(float64) < v {
				left = mid + 1
			} else {
				right = mid - 1
			}
		case float32:
			if any(orderedSlice.storage[mid]).(float32) < v {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return left
}

func (orderedSlice *OrderedSlice[T]) String() string {
	return fmt.Sprintf("%v", orderedSlice.storage)
}

func NewIntOrderedSlice(values []int) *OrderedSlice[int] {
	orderedSlice := &OrderedSlice[int]{}
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	orderedSlice.storage = append(orderedSlice.storage, values...)
	orderedSlice.length = len(orderedSlice.storage)
	return orderedSlice
}

func NewStringOrderedSlice(values []string) *OrderedSlice[string] {
	orderedSlice := &OrderedSlice[string]{}
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	orderedSlice.storage = append(orderedSlice.storage, values...)
	orderedSlice.length = len(orderedSlice.storage)
	return orderedSlice
}

func NewFloat32OrderedSlice(values []float32) *OrderedSlice[float32] {
	orderedSlice := &OrderedSlice[float32]{}
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	orderedSlice.storage = append(orderedSlice.storage, values...)
	orderedSlice.length = len(orderedSlice.storage)
	return orderedSlice
}

func NewFloat64OrderedSlice(values []float64) *OrderedSlice[float64] {
	orderedSlice := &OrderedSlice[float64]{}
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	orderedSlice.storage = append(orderedSlice.storage, values...)
	orderedSlice.length = len(orderedSlice.storage)
	return orderedSlice
}

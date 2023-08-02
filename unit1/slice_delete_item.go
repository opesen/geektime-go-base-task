package main

import (
	"errors"
	"fmt"
)

func SliceDeleteItem(slice []int, idx int) ([]int, error) {

	fmt.Printf("slice_src: %v, len=%d, cap=%d \n", slice, len(slice), cap(slice))

	if len(slice) <= idx {
		return slice, errors.New("out of range")
	}

	if idx == 0 {
		return slice[idx+1:], nil
	}

	if idx == len(slice)-1 {
		return slice[:idx], nil
	}

	slice = append(slice[:idx], slice[idx+1:]...)

	return slice, nil
}

func SliceDeleteItemPro[T any](slice []T, idx int) ([]T, error) {

	fmt.Printf("slice_src: %v, len=%d, cap=%d \n", slice, len(slice), cap(slice))

	if len(slice) <= idx {
		return slice, errors.New("out of range")
	}

	if idx == 0 {
		return slice[idx+1:], nil
	}

	if idx == len(slice)-1 {
		return slice[:idx], nil
	}

	return append(slice[:idx], slice[idx+1:]...), nil
}

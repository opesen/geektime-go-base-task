package main

import "fmt"

func main() {

	slice_src := []int{1, 3, 5, 7, 9}
	slice_dest, err := SliceDeleteItem(slice_src, 2)

	if err == nil {
		fmt.Printf("slice_dest: %v, len=%d, cap=%d \n", slice_dest, len(slice_dest), cap(slice_dest))
	}

	slice_src2 := []float64{1.5, 2.5, 3.5, 4.5}
	slice_dst2, err2 := SliceDeleteItemPro[float64](slice_src2, 0)
	if err2 == nil {
		fmt.Printf("slice_dest2 :%v, len=%d, cap=%d \n", slice_dst2, len(slice_dst2), cap(slice_dst2))
	}
}

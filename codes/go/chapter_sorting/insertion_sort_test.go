// File: insertion_sort_test.go
// Created Time: 2022-12-12
// Author: msk397 (machangxinq@gmail.com)

package chapter_sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{4, 1, 3, 1, 5, 2}
	insertionSort(nums)
	if !reflect.DeepEqual(nums,[]int{1,1,2,3,4,5}) {
		t.Errorf("insertionSort failed. Got %v, expected [1,1,2,3,4,5].", nums)
	}
	fmt.Println("插入排序完成后 nums =", nums)
}

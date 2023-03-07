package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for _, e := range arr3 {
		t.Log(e)
	}

	for idx,e := range arr3{
		t.Log(idx, e)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
    // 所有的数组元素
	arr3_sec := arr3[:]
    // 前三个元素
    arr4_sec := arr3[:3]
    // 后四个元素
    arr5_sec := arr3[1:]
    // 下面这个错误，不支持负数
    // arr6_sec := arr3[:-1]
    // 从1到3的2个元素
    arr7_sec := arr3[1:3]
	t.Log(arr3_sec, arr4_sec,arr5_sec, arr7_sec)
}

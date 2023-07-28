package main

import "fmt"

// 1、实现删除操作
// 删除切片中特定下标的元素，可以通过将该元素后面的所有元素向前移动一位，覆盖掉该元素的方式实现删除。
func deleteElement1(slice []int, index int) []int {
	if slice == nil || index < 0 || index >= len(slice) {
		return slice
	}
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

// 2、高性能实现
// 将待删除元素之后的元素向前移动n个位置，然后再通过切片的方式去掉最后一个元素。这样做的时间复杂度为O(1)，但会增加空间复杂度。
func deleteElement2(slice []int, index int) []int {
	if slice == nil || index < 0 || index >= len(slice) {
		return slice
	}
	if index < len(slice)-1 {
		copy(slice[index:], slice[index+1:])
	}
	return slice[:len(slice)-1]
}

type Number interface {
	int | uint
}

// 3、泛型方法实现
func deleteElement3[E Number](slice []E, index int) []E {
	if slice == nil || index < 0 || index >= len(slice) {
		return slice
	}
	if index < len(slice)-1 {
		copy(slice[index:], slice[index+1:])
	}
	return slice[:len(slice)-1]
}

// 4、支持缩容
// 在删除元素之后判断切片的长度是否小于其容量的一半，如果是，则将其容量缩小一半。
func deleteElement4[E Number](slice []E, index int) []E {
	if slice == nil || index < 0 || index >= len(slice) {
		return slice
	}
	if index < len(slice)-1 {
		copy(slice[index:], slice[index+1:])
	}
	slice = slice[:len(slice)-1]
	defaultCap := 2
	if len(slice) < cap(slice)/2 {
		newCap := cap(slice) / 2
		if newCap < defaultCap {
			newCap = defaultCap
		}
		newSlice := make([]E, len(slice), newCap)
		copy(newSlice, slice)
		slice = newSlice
	}
	return slice
}

func main() {
	slice0 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice0[:])

	//slice1 := deleteElement1(slice0, 2)
	//fmt.Println(slice1[:])

	//slice2 := deleteElement2(slice0, 2)
	//fmt.Println(slice2[:])

	//slice3 := deleteElement3(slice0, 2)
	//fmt.Println(slice3[:])

	slice4 := deleteElement4(slice0, 2)
	fmt.Println(slice4[:])
}

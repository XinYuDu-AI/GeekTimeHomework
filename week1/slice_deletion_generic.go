package main

import "fmt"

// 使用泛型实现删除切片中指定位置的元素，若成功删除返回修改后的切片和一个布尔值
func RemoveAtIndex[T any](idx int, s []T) ([]T, error) {
	if idx < 0 || idx >= len(s) {
		return s, fmt.Errorf("index out of range")
	}
	// fmt.Printf("Before removal, len(s) =%v cap(s) =%v\n", len(s), cap(s))
	// new := append(s[:idx], s[idx+1:]...)
	// fmt.Printf("After removal, len(s) =%v cap(s) =%v\n", len(new), cap(new))

	return append(s[:idx], s[idx+1:]...), nil
}

// 删除切片中指定位置的元素，若成功删除返回修改后的切片和一个布尔值
// 当切片实际使用长度小于等于其容量的一半时，缩容
func RemoveAtIndexThenShrink[T any](idx int, s []T) ([]T, error) {
	if idx < 0 || idx >= len(s) {
		return s, fmt.Errorf("index out of range")
	}
	// 删除元素后切片实际使用空间大于原切片的一半时，不需要缩容
	if (len(s)-1)*2 > cap(s) {
		return RemoveAtIndex(idx, s)
	}
	// 需要缩小容量
	new := make([]T, len(s)-1)
	copy(new, s[:idx])
	copy(new[idx:], s[idx+1:])
	return new, nil
}

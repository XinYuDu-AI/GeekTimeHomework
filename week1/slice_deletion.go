package main

import "fmt"

// V0版本的手动实现法：删除切片中指定位置的元素，若成功删除返回修改后的切片和一个布尔值
// true代表成功删除，false代表删除失败
func RemoveAtIndexV0(idx int, s []int) ([]int, error) {
	if idx < 0 || idx >= len(s) {
		return s, fmt.Errorf("index out of range")
	}
	for i := idx; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	// 这里做切片的原因是，复制最后一个元素到前一个位置的时候
	// 最后位置的元素还在，所以需要切片去掉最后一个元素
	return s[:len(s)-1], nil
}

// V1版本的内置函数法：删除切片中指定位置的元素，若成功删除返回修改后的切片和一个布尔值
func RemoveAtIndexV1(idx int, s []int) ([]int, error) {
	if idx < 0 || idx >= len(s) {
		return s, fmt.Errorf("index out of range")
	}
	return append(s[:idx], s[idx+1:]...), nil
}

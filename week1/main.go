package main

import "fmt"

func main() {
	fmt.Println("---------Using V0:---------")
	a := []int{1, 2, 3, 4, 5}
	if newA, ok := RemoveAtIndexV0(2, a); ok != nil {
		fmt.Println("After removal, a =", newA)
	} else {
		fmt.Println("Removal was unsuccessful")
	}
	fmt.Println("Original a =", a)

	fmt.Println("\n---------Using V1:---------")
	a = []int{1, 2, 3, 4, 5}
	if newA, ok := RemoveAtIndexV1(2, a); ok != nil {
		fmt.Println("After removal, a =", newA)
	} else {
		fmt.Println("Removal was unsuccessful")
	}
	fmt.Println("Original a =", a)

	fmt.Println("\n---------Using generic:---------")
	a = []int{1, 2, 3, 4, 5}
	if newA, ok := RemoveAtIndex(2, a); ok != nil {
		fmt.Println("After removal, a =", newA)
	} else {
		fmt.Println("Removal was unsuccessful")
	}
	fmt.Println("Original a =", a)

	fmt.Println("\n---------Using generic with shrink:---------")
	a = []int{1, 2, 3, 4, 5, 6}

	newA, err := RemoveAtIndexThenShrink(5, a)
	if err != nil {
		fmt.Printf("After removal, len(a) =%v, cap(a)=%v\n", len(newA), cap(newA))
	}

	newA, err = RemoveAtIndexThenShrink(3, a)
	if err != nil {
		fmt.Printf("After removal, len(a) =%v, cap(a)=%v\n", len(newA), cap(newA))
	}

	newA, err = RemoveAtIndexThenShrink(1, a)
	if err != nil {
		fmt.Printf("After removal, len(a) =%v, cap(a)=%v\n", len(newA), cap(newA))
	}

	fmt.Println("Original a =", a)
}

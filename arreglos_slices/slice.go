package arreglosslices

import (
	"fmt"
)

var tabla []int = []int{2, 5, 4}

var arreglo [10]int = [10]int{6, 98, 5, 7, 998, 6, 8, 6, 7, 41}

func MostrarSlice() {
	fmt.Println(tabla)

	parte := arreglo[3:]
	parteDos := arreglo[:3]
	fmt.Println(parte)
	fmt.Println(parteDos)
}

func Capacidad() {
	slice := make([]int, 5, 20)
	fmt.Printf("Large %d, Capacidad %d \n", len(slice), cap(slice))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Large %d, Capacidad %d", len(nums), cap(nums))
}

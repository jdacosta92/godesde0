package arreglos_slices

import (
	"fmt"
)

var tabla_slice []int = []int{2, 5, 4}

func MuestroSlice() {
	fmt.Println(tabla_slice)
}

func Capacidad() {
	//elementos := make([]int, 5, 20)
	//fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
	fmt.Println()
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}

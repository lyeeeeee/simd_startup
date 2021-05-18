package main

import (
	"fmt"
	"time"
	"unsafe"
)

func SumAssem(p unsafe.Pointer, size int) int
func SumAssem4Way(p unsafe.Pointer, size int) int
func SumAssemSIMD(acc *[4]uint64, p unsafe.Pointer, size int) int

const size = 1<<20

func main() {
	dat0 := [size]int{}

	for i := 0; i < size; i++ {
		dat0[i] = i
	}

	fmt.Println(size)

	start := time.Now()
	res0 := SumArray(unsafe.Pointer(&dat0),size)
	duration0 := time.Since(start).Nanoseconds()
	fmt.Println("method0", duration0, "ns", res0)

	start = time.Now()
	res1 := SumScalar(unsafe.Pointer(&dat0), size)
	duration1 := time.Since(start).Nanoseconds()
	fmt.Println("method1", duration1, "ns", res1)

	start = time.Now()
	res2 := SumAssem(unsafe.Pointer(&dat0), size)
	duration2 := time.Since(start).Nanoseconds()
	fmt.Println("method2", duration2, "ns", res2)

	start = time.Now()
	res3 := SumAssem4Way(unsafe.Pointer(&dat0), size)
	duration3 := time.Since(start).Nanoseconds()
	fmt.Println("method3", duration3, "ns", res3)

	start = time.Now()
	accs := [4]uint64{0,0,0,0}

	SumAssemSIMD(&accs, unsafe.Pointer(&dat0), size)
	res4 := uint64(0)
	for i := 0; i < 4; i++ {
		res4 += accs[i]
	}

	duration4 :=time.Since(start).Nanoseconds()
	fmt.Println("method4", duration4, "ns", res4)
}

func SumArray(p unsafe.Pointer, size int) int {
	sum := 0
	for i := 0; i < size; i += 1 {
		sum += *(*int)(p)
		p = unsafe.Pointer(uintptr(p) + 8)
	}
	return sum
}

func SumScalar(p unsafe.Pointer, size int) int {
	sum := 0
	for i := 0; i < size; i += 4 {
		sum += *(*int)(p)
		sum += *(*int)(unsafe.Pointer(uintptr(p) + 8))
		sum += *(*int)(unsafe.Pointer(uintptr(p) + 16))
		sum += *(*int)(unsafe.Pointer(uintptr(p) + 24))
		p = unsafe.Pointer(uintptr(p) + 32)
	}

	return sum
}

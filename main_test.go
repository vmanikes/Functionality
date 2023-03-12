// Package main contains {ENTER PACKAGE DESCRIPTION}
package functionality

import (
	"fmt"
	"testing"
)

func TestDsd(t *testing.T) {
	val := Iterator[int]{1, 2, 3, 4}

	fmt.Printf("Address of array = %v: %p\n", val, &val)

	newVal := val.Filter(func(i int) bool {
		return i%2 == 0
	})

	fmt.Println(val, newVal)

	fmt.Printf("Address of array = %v: %p\n", newVal, &newVal)
}

//func BenchmarkName(b *testing.B) {
//	b.ReportAllocs()
//
//	val := make(Iterator[int], 0, 1000)
//
//	for i := 0; i < 1000; i++ {
//		val = append(val, i)
//	}
//
//	for i := 0; i < b.N; i++ {
//		val = val.filter(func(i int) bool {
//			return i%2 == 0
//		})
//	}
//}

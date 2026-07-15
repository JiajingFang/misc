package data_structure

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

var num = 1000


//https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
func BenchmarkHashmap(b *testing.B) {
	size := 1000
	hm := NewHashMap(&size)
	for j:= 0; j < size; j++ {
		//put 1000 keyValue pairs into Hashmap
		value := strconv.Itoa(rand.Intn(1000) + 1)
		hm.Set(strconv.Itoa(j), value)
	}
	fmt.Println("hash map benchmark results for get value")
    for i := 0; i < b.N; i++ {
		hm.Get(strconv.Itoa(rand.Intn(1000)))
    }
}

func BenchmarkSlices(b *testing.B) {
	size := 1000
	numbers := []string{}
	for j:= 0; j < size; j++ {
		//put 1000 elements into slice
		numbers = append(numbers, strconv.Itoa(rand.Intn(1000) + 1))
	}
	fmt.Println("slices benchmark results for get value")
    for i := 0; i < b.N; i++ {
		_ = slices.Index(numbers, strconv.Itoa(rand.Intn(1000)))
    }

}
package data_structure

import (
	"math/rand"
	"slices"
	"strconv"
	"testing"
)

var num = 1000


//https://blog.logrocket.com/benchmarking-golang-improve-function-performance/
func BenchmarkHashmap(b *testing.B) {
	size := 1000000
	hm := NewHashMap(&size)
	for j:= 0; j < size; j++ {
		//put 1000 keyValue pairs into Hashmap
		value := strconv.Itoa(rand.Intn(1000) + 1)
		hm.Set(strconv.Itoa(j), value)
	}
    for i := 0; i < b.N; i++ {
		hm.Get(strconv.Itoa(rand.Intn(1000)))
    }
}

func BenchmarkRealHashmap(b *testing.B) {
	size := 1000000
	hm := make(map[string]string, size)
	for j:= 0; j < size; j++ {
		//put 1000 keyValue pairs into Hashmap
		value := strconv.Itoa(rand.Intn(1000) + 1)
		hm[strconv.Itoa(j)] = value
	}
    for i := 0; i < b.N; i++ {
		_ = hm[strconv.Itoa(rand.Intn(1000))]
    }
}

func BenchmarkSlices(b *testing.B) {
	size := 1000000
	numbers := []string{}
	for j:= 0; j < size; j++ {
		//put 1000 elements into slice
		numbers = append(numbers, strconv.Itoa(rand.Intn(1000) + 1))
	}
    for i := 0; i < b.N; i++ {
		_ = slices.Index(numbers, strconv.Itoa(rand.Intn(1000)))
    }

}
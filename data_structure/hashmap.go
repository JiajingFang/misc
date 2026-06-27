package data_structure

import (
	"hash/fnv"
	"math"
)

const defaultBucketSize = 16

type HashMap struct {
	buckets [][]keyValuePair
	bucketSize int
}

type keyValuePair struct {
	key string
	value string
}

func NewHashMap(size *int) *HashMap {
	if size == nil {
		return &HashMap{
			buckets: make([][]keyValuePair, defaultBucketSize),
			bucketSize: defaultBucketSize,
		}
	} else {
		hSize := 1 << int(math.Ceil(math.Log2(float64(*size))))
		return &HashMap{
			buckets: make([][]keyValuePair, hSize),
			bucketSize: hSize,
		}
	}
}

func (h *HashMap) getHash (key string) uint32 {
	// get the HashValue from hash/fnv then apply Hash Disturbance Function
	hV := fnv.New32a()
	hV.Write([]byte(key))
	return hV.Sum32() ^ (hV.Sum32() >> 16)
}

func (h *HashMap) getIndex (hash uint32, length uint32) uint32 {
	return (length - 1) & hash
}

func (h *HashMap) Set (key, value string)  {
	// get hash value, bitwise AND op with bucket size to get hash mapping index
	// if bucket is empty, add; if key exits, replace; if key not exists, append

	hV := h.getHash(key)
	id := h.getIndex(hV, uint32(h.bucketSize))

	exist := false
	for _, kvPair := range h.buckets[id] {
		if kvPair.key == key {
			kvPair.value = value
			exist = true
			break
		}
	}
	if !exist {
		h.buckets[id] = append(h.buckets[id], keyValuePair{key: key, value: value})
	}
}

func (h *HashMap) Get (key string) (*string) {
	// get hash value, bitwise AND op with bucket size to get hash mapping index
	// traverse the bucket, if key matches, return value; if key doesn't matches, return nil

	hV := h.getHash(key)
	id := h.getIndex(hV, uint32(h.bucketSize))

	for _, kvPair := range h.buckets[id] {
		if kvPair.key == key {
			return &kvPair.value
		}
	}
	return nil

}
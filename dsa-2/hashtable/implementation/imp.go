package main

import "fmt"

type KeyValue struct {
	Key   string
	Value interface{}
}

type HashTable struct {
	Buckets    [][]KeyValue
	BucketSize int
}

func (ht *HashTable) Hash(key string) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % ht.BucketSize
}

func (ht *HashTable) Put(key string, value interface{}) {
	index := ht.Hash(key)
	for i, kv := range ht.Buckets[index] {
		if kv.Key == key {
			ht.Buckets[index][i].Value = value
		}
	}
	ht.Buckets[index] = append(ht.Buckets[index], KeyValue{Key: key, Value: value})
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	index := ht.Hash(key)
	for _, kv := range ht.Buckets[index] {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return nil, false
}

func (ht *HashTable) Delete(key string) bool {
	index := ht.Hash(key)
	for i, kv := range ht.Buckets[index] {
		if kv.Key == key {
			ht.Buckets[index] = append(ht.Buckets[index][:i], ht.Buckets[index][i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	hashTable := HashTable{Buckets: make([][]KeyValue, 5), BucketSize: 5}

	// Add keys that may hash to the same bucket
	hashTable.Put("abc", "Value1")
	hashTable.Put("cab", "Value2") // Potential collision with "abc"

	// Retrieve and print values
	if value, found := hashTable.Get("abc"); found {
		fmt.Println("abc:", value)
	}
	if value, found := hashTable.Get("cab"); found {
		fmt.Println("cab:", value)
	}
}

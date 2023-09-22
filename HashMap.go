package main

import (
	"fmt"
)

type MyHashMap struct {
	keys      []int
	values    []int
	size      int
	curr_size int
	loadlimit int
}

func Constructor() MyHashMap {

	size := 5
	keys := make([]int, size)
	values := make([]int, size)
	loadlimit := size - (size / 10)
	for i := 0; i < size; i++ {
		keys[i] = -1
		values[i] = -1
	}
	return MyHashMap{
		keys:      keys,
		values:    values,
		size:      size,
		curr_size: 0,
		loadlimit: loadlimit,
	}

}

func (this *MyHashMap) Resize() {
	new_size := 2 * this.size
	new_keys := make([]int, new_size)
	new_values := make([]int, new_size)
	for i := 0; i < new_size; i++ {
		new_keys[i] = -1
		new_values[i] = -1
	}
	for index, key := range this.keys {
		if key == -1 {
			continue
		}
		new_index := GetHash(key, new_size, new_keys)
		new_keys[new_index] = key
		new_values[new_index] = this.values[index]
	}
	this.keys = new_keys
	this.values = new_values
	this.loadlimit = new_size - (new_size / 10)
	this.size = new_size
}

func GetHash(key int, size int, keys []int) int {
	n := key % size
	for keys[n] != -1 && keys[n] != key {
		n = (n + 1) % size
	}
	return n
}

func (this *MyHashMap) Put(key int, value int) {
	//get the hash

	if this.curr_size == this.loadlimit {
		this.Resize()
		fmt.Println("resize finished", this.keys)
	}
	index := GetHash(key, this.size, this.keys)
	this.keys[index] = key
	this.values[index] = value
	this.curr_size += 1
}

func (this *MyHashMap) Get(key int) int {
	index := GetHash(key, this.size, this.keys)
	return this.values[index]
}

func (this *MyHashMap) Remove(key int) {
	index := GetHash(key, this.size, this.keys)
	if this.keys[index] != -1 {
		this.curr_size -= 1
	}
	this.keys[index] = -1
	this.values[index] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 */
func main() {
	obj := Constructor()
	for i := 0; i < 15; i++ {
		obj.Put(i*2, i)
	}
	param_2 := obj.Get(1)
	fmt.Println(param_2)
	//obj.Remove(1)
	fmt.Println(obj.keys)
}

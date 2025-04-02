package dsa_01_12_25

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	arr  []int
	size int
	set  map[int]int
}

func Constructor() RandomizedSet {
	arr := make([]int, 0)
	size := 0
	set := make(map[int]int)
	return RandomizedSet{
		arr:  arr,
		size: size,
		set:  set,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}

	this.arr = append(this.arr, val)
	this.set[val] = this.size
	this.size++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.set[val]
	if !ok {
		return false
	}

	// Swap the element to remove with the last element
	lastElement := this.arr[this.size-1]
	fmt.Println(lastElement, index, this.arr)
	this.arr[index] = lastElement
	this.set[lastElement] = index
	fmt.Println(this.set, this.arr)

	// Remove the last element
	this.arr = this.arr[:this.size-1]
	delete(this.set, val)
	this.size--
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(this.size)
	return this.arr[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func Main() {

	randSet := Constructor()
	randSet.Insert(2)
	randSet.Insert(1)
	randSet.Remove(2)
	// randSet.Insert(2)
	// randSet.GetRandom()
	// randSet.Remove(1)
	// randSet.Insert(2)
	// randSet.GetRandom()
}

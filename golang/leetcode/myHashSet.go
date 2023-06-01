package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	obj.Add(3)
	obj.Add(10)

	obj.Remove(2)

	fmt.Println(obj.Contains(2))
	fmt.Println(obj.Contains(1))
}

type MyHashSet struct {
	List map[int]bool
}

func Constructor() MyHashSet {
	list := map[int]bool{}

	return MyHashSet{
		List: list,
	}
}

func (this *MyHashSet) Add(key int) {
	this.List[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.List[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	if val, ok := this.List[key]; ok && val {
		return true
	} else {
		return false
	}
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

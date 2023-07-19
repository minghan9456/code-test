package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	obj := Constructor(3)

	obj.Set(0, 5)
	obj.Set(1, 2)
	fmt.Println("snap", obj.Snap())
	// fmt.Println("get", obj.Get(0, 0))
	// fmt.Println("get", obj.Get(1, 0))
	fmt.Println("snap", obj.Snap())
	obj.Set(1, 3)
	fmt.Println("snap", obj.Snap())
	fmt.Println("snap", obj.Snap())
	fmt.Println("snap", obj.Snap())
	fmt.Println("snap", obj.Snap())
	// fmt.Println("get", obj.Get(0, 0))
	// fmt.Println("get", obj.Get(1, 0))
	fmt.Println("get", obj.Get(1, 1))
	fmt.Println("get", obj.Get(1, 2))
	/*
		fmt.Println(obj.Get(0, 0))
		fmt.Println(obj.Get(1, 0))
		fmt.Println(obj.Get(1, 1))
	*/
}

type Elem struct {
	val     int
	version int
}

type SnapshotArray struct {
	History [][]Elem
	SnapID  int
}

func Constructor(length int) SnapshotArray {
	hs := make([][]Elem, length)

	return SnapshotArray{
		History: hs,
		SnapID:  0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	// Update value if we are at the same version
	if len(this.History[index]) > 0 {
		tmp := this.History[index]

		if tmp[len(tmp)-1].version == this.SnapID {
			tmp[len(tmp)-1].val = val
			return
		}
	}

	// Add new record if we are at the higher version
	this.History[index] = append(this.History[index], Elem{val, this.SnapID})
}

func (this *SnapshotArray) Snap() int {
	curr := this.SnapID
	this.SnapID++
	return curr
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	// Binary search for the exact snap_id or the first below it
	arr := this.History[index]

	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid].version == snap_id {
			return arr[mid].val
		} else if arr[mid].version < snap_id {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left-1 <= len(arr) && left-1 >= 0 {
		return arr[left-1].val
	}

	return 0
}

/**
// out of memory
type SnapshotArray struct {
	CurrSnap int
	Array    []int
	Cache    [][]int
}

func Constructor(length int) SnapshotArray {
	arr := make([]int, length)
	cache := [][]int{}

	return SnapshotArray{
		CurrSnap: 0,
		Array:    arr,
		Cache:    cache,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.Array[index] = val
}

func (this *SnapshotArray) Snap() int {
	curr := this.CurrSnap

	dst := make([]int, len(this.Array))
	copy(dst, this.Array)
	this.Cache = append(this.Cache, dst)

	this.CurrSnap++
	return curr
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	return this.Cache[snap_id][index]
}

 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
*/

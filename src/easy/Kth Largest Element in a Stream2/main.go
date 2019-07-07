package main

import (
	"fmt"
)

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, arr)
	var length int
	length = kthLargest.Add(3) // returns 4
	fmt.Println(length)
	length = kthLargest.Add(5) // returns 5
	fmt.Println(length)
	length = kthLargest.Add(10) // returns 5
	fmt.Println(length)
	length = kthLargest.Add(9) // returns 8
	fmt.Println(length)
	length = kthLargest.Add(4) // returns 8
	fmt.Println(length)
}

type Heap struct {
	buf []int
	sz  int
	k   int
}

func (h *Heap) GetKth() int {
	for len(h.buf) > h.k {
		h.Extract()
	}

	return h.buf[h.k-1]
}

func (h *Heap) Extract() {
	x := h.buf[len(h.buf)-1]
	index := 0
	for index*2+1 < h.sz {
		a := index*2 + 1
		b := index*2 + 2
		if b < h.sz && h.buf[b] < h.buf[a] {
			a = b
		}

		if h.buf[a] >= x {
			break
		}
		h.buf[index] = h.buf[a]
		index = a
	}
}

func (h *Heap) Push(key int) int {
	/*
		自分のノードの番号（一番最初は一番右に追加する予定）
		最後のノードに追加する
	*/
	h.buf = append(h.buf, key)
	index := len(h.buf) - 1
	h.sz = index
	h.sz++
	for index > 0 {
		//親のノード番号
		parent := (index - 1) / 2
		//親の値と等しくなったら
		if h.buf[parent] <= key {
			break
		}
		//親のノードの数字を下ろして自分が上に行く
		h.buf[parent], h.buf[index] = h.buf[index], h.buf[parent]
		index = parent
	}
	return index
}

type KthLargest struct {
	heap Heap
}

func Constructor(k int, nums []int) KthLargest {
	heap := Heap{}
	heap.sz = 0
	heap.k = k
	heap.buf = make([]int, 0)
	for _, Val := range nums {
		heap.Push(Val)
	}
	kthLargest := KthLargest{}
	kthLargest.heap = heap
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	this.heap.Push(val)
	return this.heap.GetKth()
}

/*
TreeNode
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
ListNode
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//Reverse string reverse
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//SampleStruct privateで構造体宣言
type SampleStruct struct {
	Index  int
	Height int
}

/*
sort用
GO 1.6ではsort#Sliceが使えないのでこちらで対応
*/
type samples []SampleStruct

func (u samples) Len() int {
	return len(u)
}

func (u samples) Less(i, j int) bool {
	return u[i].Height > u[j].Height
}

func (u samples) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

//Int64Abs abstract int64
func Int64Abs(a int64) int64 {
	if a < 0 {
		a *= -1
	}
	return a
}

//IntAbs abstract int
func IntAbs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

//IntMin return minimum value
func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//IntMax return maximum value
func IntMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

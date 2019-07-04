package main

import (
	"fmt"
	"sort"
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

type KthLargest struct {
	Nums []int
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	obj := new(KthLargest)
	sort.Sort(sort.IntSlice(nums))
	obj.Nums = nums
	obj.K = k
	return *obj
}

func (this *KthLargest) Add(val int) int {
	i := sort.Search(len(this.Nums), func(i int) bool { return this.Nums[i] >= val })
	if i < len(this.Nums) {
		this.Nums = append(this.Nums[:i+1], this.Nums[i:]...)
		this.Nums[i] = val
	} else {
		this.Nums = append(this.Nums, val)
	}
	return this.Nums[len(this.Nums)-1-this.K+1]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
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

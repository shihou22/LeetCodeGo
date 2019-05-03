package main

import "sort"

func main() {

}

/*
ListNode
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoListsRecursion(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var array = make([]int, 0)
	for l1 != nil || l2 != nil {
		if l1 != nil {
			array = append(array, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			array = append(array, l2.Val)
			l2 = l2.Next
		}
	}
	sort.Sort(sort.IntSlice(array))

	var res = ListNode{0, nil}
	var wk *ListNode
	wk = &res
	for index := 0; index < len(array); index++ {
		var tmp = ListNode{array[index], nil}
		wk.Next = &tmp
		wk = wk.Next
	}
	return res.Next
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

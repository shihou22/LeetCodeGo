package main

import (
	"sort"
)

func main() {

}

func minMoves(nums []int) int {

	if nums == nil || len(nums) == 1 {
		return 0
	}

	sort.Sort(sort.IntSlice(nums))
	incBase := len(nums) - 1
	maxBase := nums[incBase]
	minBase := nums[0]
	total := 0
	for _, val := range nums {
		total += maxBase - val
	}

	needMin := total - (maxBase-minBase)*incBase
	if needMin < 0 {
		needMin *= -1
	}

	res := 0
	if total-needMin == 0 {
		res = total / incBase
	} else {
		res = (maxBase - minBase) + needMin //lmd(needMin, incBase)/incBase
	}
	return res
}

func lmd(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	wkA := a
	wkB := b
	if a < b {
		tmp := wkA
		wkA = wkB
		wkB = tmp
	}
	amari := wkA % wkB
	for amari > 0 {
		wkA = wkB
		wkB = amari
		amari = wkA % wkB
	}
	return wkB
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

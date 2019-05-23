package main

import "sort"

func main() {

}
func containsDuplicate(nums []int) bool {

	memo := make(map[int]bool)
	for index := 0; index < len(nums); index++ {
		if memo[nums[index]] {
			return true
		} else {
			memo[nums[index]] = true
		}
	}

	return false

}

func containsDuplicateUseSort(nums []int) bool {

	if len(nums) < 2 {
		return false
	}

	sort.Sort(sort.IntSlice(nums))

	for index := 1; index < len(nums); index++ {
		if nums[index-1] == nums[index] {
			return true
		}
	}

	return false

}

func containsDuplicateUseMap(nums []int) bool {

	memo := make(map[int]bool)
	for index := 0; index < len(nums); index++ {
		if _, ok := memo[nums[index]]; ok {
			return true
		} else {
			memo[nums[index]] = true
		}
	}

	return false

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

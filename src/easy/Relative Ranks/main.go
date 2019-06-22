package main

import (
	"sort"
	"strconv"
)

func main() {

}

func findRelativeRanks(nums []int) []string {

	cop := append([]int{}, nums...)
	memo := make(map[int]string)
	sort.Sort(sort.IntSlice(cop))
	cnt := 1
	for index := len(cop) - 1; index >= 0; index-- {
		switch cnt {
		case 1:
			memo[cop[index]] = "Gold Medal"
		case 2:
			memo[cop[index]] = "Silver Medal"
		case 3:
			memo[cop[index]] = "Bronze Medal"
		default:
			memo[cop[index]] = strconv.Itoa(cnt)
		}
		cnt++
	}

	res := make([]string, 0)
	for _, val := range nums {
		res = append(res, memo[val])
	}
	return res
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

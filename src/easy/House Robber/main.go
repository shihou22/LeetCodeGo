package main

import "sort"

func main() {

}

/*
it's good problem
*/
func rob(nums []int) int {
	rob, norob := 0, 0
	for _, num := range nums {
		rob, norob = num+norob, max(norob, rob)
		// var robT = num + norob
		// var norobT = max(norob, rob)
		// rob = robT
		// norob = norobT
	}
	return max(norob, rob)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func robDP1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var wk = make([]int, len(nums))
	wk[0] = nums[0]
	wk[1] = nums[1]
	for index := range nums {
		for inner := index + 2; inner < len(nums); inner++ {
			if wk[inner] < wk[index]+nums[inner] {
				wk[inner] = wk[index] + nums[inner]
			}
		}
	}
	sort.Sort(sort.IntSlice(wk))
	return wk[len(nums)-1]
}

func robTLE(nums []int) int {
	return rescursiveRob(nums, 0)
}
func rescursiveRob(nums []int, currentI int) int {

	if currentI >= len(nums) {
		return 0
	}
	var cur1 = rescursiveRob(nums, currentI+2) + nums[currentI]
	var cur2 = rescursiveRob(nums, currentI+1)
	if cur1 > cur2 {
		return cur1
	}
	return cur2
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

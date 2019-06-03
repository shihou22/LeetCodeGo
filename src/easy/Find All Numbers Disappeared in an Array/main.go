package main

import "sort"

func main() {

}

func findDisappearedNumbers(nums []int) []int {

	for index := 0; index < len(nums); index++ {
		if nums[index] != index+1 {
			for nums[index] != index+1 && nums[nums[index]-1] != nums[index] {
				nums[index], nums[nums[index]-1] = nums[nums[index]-1], nums[index]
			}
		}
	}

	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

func findDisappearedNumbersSlow(nums []int) []int {
	sort.Sort(sort.IntSlice(nums))
	res := make([]int, 0)
	max := len(nums)
	left := 1
	right := 0
	for left <= max {

		isExit := false
		for right < len(nums) && left == nums[right] {
			isExit = true
			right++
		}
		if !isExit {
			res = append(res, left)
		}
		left++
	}

	return res
}

func findDisappearedNumbersOutOfRange(nums []int) []int {
	res := make([]int, 0)
	key := 0
	for index := 1; index <= len(nums); index++ {
		key = int(uint(key) << uint(1))
		key++
	}
	for _, val := range nums {
		index := int(uint64(1) << uint64(val-1))
		if (key & index) > 0 {
			key ^= index
		}
	}
	for cnt := 0; cnt < len(nums); cnt++ {
		index := uint64(1) << uint64(cnt)
		if (key & int(index)) > 0 {
			res = append(res, cnt+1)
		}
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

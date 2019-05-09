package main

import "math"

func main() {

}

/*
ListNode
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
[-2,1,-3,4,-1,2,1,-5,4]
*/
func maxSubArray(nums []int) int {
	var res = nums[0]
	var keepMax = nums[0]
	for index := 1; index < len(nums); index++ {
		/*
			-になったらreset扱い
		*/
		if 0 > keepMax {
			keepMax = 0
		}
		keepMax += nums[index]
		res = int(math.Max(float64(res), float64(keepMax)))
	}
	// var res = math.MinInt32
	// for left := 0; left < len(nums); left++ {
	// 	var wk = nums[left]
	// 	res = int(math.Max(float64(res), float64(wk)))
	// 	for rigth := left + 1; rigth < len(nums); rigth++ {
	// 		wk += nums[rigth]
	// 		res = int(math.Max(float64(res), float64(wk)))
	// 	}

	// }

	return res
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

package main

import "fmt"

func main() {

}

func moveZeroesOther2(nums []int) {

	var zero = 0
	for _, val := range nums {
		if val == 0 {
			zero++
		}
	}
	var res = make([]int, 0)
	for _, val := range nums {
		if val != 0 {
			res = append(res, val)
		}
	}
	for index := 0; index < zero; index++ {
		res = append(res, 0)
	}
	for Index, val := range res {
		nums[Index] = val
	}
	fmt.Println("hoge")
}

func moveZeroesOther(nums []int) {

	var start = 0
	var end = 0
	for start < len(nums) && end < len(nums) {
		for start < len(nums) && nums[start] != 0 {
			start++
		}
		end = start + 1
		for end < len(nums) && nums[end] == 0 {
			end++
		}
		if start <= end && end < len(nums) && nums[end] != 0 {
			nums[start], nums[end] = nums[end], nums[start]
		}
	}
	fmt.Println("hoge")
}

func moveZeroes(nums []int) {

	for index, element := range nums {

		if element == 0 {
			for inner := index + 1; inner < len(nums); inner++ {
				if nums[inner] != 0 {
					nums[index], nums[inner] = nums[inner], nums[index]
					break
				}
			}
		}

	}
	// fmt.Println("hoge")
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

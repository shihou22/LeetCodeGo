package main

import (
	"fmt"
	"strconv"
)

func main() {

	for index := 1; index < 200; index++ {
		res := strconv.Itoa(index) + " / " + strconv.Itoa(findNthDigit(index))
		fmt.Println(res)
	}

}

func findNthDigit(n int) int {
	digitAmount := 9
	start := 1
	length := 1
	for n > length*digitAmount {
		n -= length * digitAmount
		digitAmount *= 10
		start *= 10
		length++
	}
	offset := (n - 1) / length
	start += offset

	return start % 10
}

func pow(n int) int {
	m := 1
	for i := 0; i < n; i++ {
		m *= 10
	}

	return m
}

func get(n, m int) int {
	s := strconv.Itoa(n)
	return int(s[m] - '0')
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

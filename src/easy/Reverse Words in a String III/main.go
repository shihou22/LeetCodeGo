package main

import "strings"

func main() {

}

func reverseWords(s string) string {

	tmp := strings.Split(s, " ")

	for index := 0; index < len(tmp); index++ {
		tmp[index] = Reverse(tmp[index])
	}
	return strings.Join(tmp, " ")
}

//Reverse string reverse
func Reverse(s string) string {
	wk := []rune(s)
	for i, j := 0, len(wk)-1; i < j; i, j = i+1, j-1 {
		wk[i], wk[j] = wk[j], wk[i]
	}
	return string(wk)
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

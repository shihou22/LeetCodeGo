package main

import "fmt"

func main() {

}

func checkSlice() {
	var stack = make([]int, 0)
	for index := 0; index < 10; index++ {
		stack = append(stack, index)
	}

	var wkStack = make([]int, len(stack))
	copy(wkStack, stack)

	copy(stack[2:], stack[3:])
	stack = stack[:len(stack)-1]

	wkStack[2] = wkStack[len(wkStack)-1]
	wkStack = wkStack[:len(wkStack)-1]
	fmt.Println(stack)
}
func isPalindrome(head *ListNode) bool {

	if head == nil {
		return true
	}
	var stack = make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	var lengthStack = len(stack)
	for index := 0; index < lengthStack/2; index++ {
		if stack[index] != stack[lengthStack-1-index] {
			return false
		}
	}

	return true
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

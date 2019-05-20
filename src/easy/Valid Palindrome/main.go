package main

import "strings"

func main() {

}

func isValid(b byte) bool {
	if ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {

	if s == "" || len(s) == 1 {
		return true
	}
	start := 0
	end := len(s) - 1
	wk := strings.ToUpper(s)
	unStrCnt := 0
	for start < end {
		for start < len(s)-1 && !isValid(wk[start]) {
			start++
			unStrCnt++
		}
		for end > 0 && !isValid(wk[end]) {
			end--
			unStrCnt++
		}
		if wk[start] == wk[end] {
			start++
			end--
		} else {
			if unStrCnt == len(s) {
				return true
			} else {
				return false
			}
		}
	}

	return true
}
func isPalindromeOther(s string) bool {

	if s == "" {
		return true
	}
	start := 0
	end := len(s) - 1
	wk := strings.ToUpper(s)
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	unStrCnt := 0
	for start < end {
		front := wk[start : start+1]
		back := wk[end : end+1]
		for start < len(s)-1 && !strings.Contains(str, front) {
			start++
			unStrCnt++
			front = wk[start : start+1]
		}
		for end > 0 && !strings.Contains(str, back) {
			end--
			unStrCnt++
			back = wk[end : end+1]
		}
		if front == back {
			start++
			end--
		} else {
			if unStrCnt == len(s) {
				return true
			} else {
				return false
			}
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

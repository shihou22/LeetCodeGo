package main

import (
	"strings"
)

func main() {

}

func licenseKeyFormatting(S string, K int) string {

	res := make([]byte, 0)
	wk := strings.ToUpper(S)
	cnt := 0
	for index := len(wk) - 1; index >= 0; index-- {
		if S[index] != '-' {
			cnt++
			if cnt > K {
				res = append(res, '-')
				cnt = 1
			}
			res = append(res, wk[index])

		}
	}
	//reverse array
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

//Reverse string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func licenseKeyFormattingOld(S string, K int) string {

	res := ""
	wk := strings.ToUpper(S)
	cnt := 0
	for index := len(wk) - 1; index >= 0; index-- {
		if S[index] != '-' {
			cnt++
			if cnt > K {
				res = "-" + res
				cnt = 1
			}
			res = string(wk[index]) + res

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

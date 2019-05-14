package main

import "strings"

func main() {

}

func repeatedSubstringPractice(s string) bool {

	var sLen = len(s)
	if sLen <= 1 {
		return false
	}

	for index := len(s) / 2; index > 0; index-- {
		var wk = s[0:index]
		var wkRes string
		for inner := 0; inner < len(s)/index; inner++ {
			wkRes += wk
		}
		if wkRes == s {
			return true
		}
	}

	return false
}

func repeatedSubstringPatternOther(s string) bool {

	var sLen = len(s)
	if sLen <= 1 {
		return false
	}

	for index := sLen / 2; index > 0; index-- {

		var repS = s[0:index]
		var sForJudge string
		for inner := 0; inner < sLen/index; inner++ {
			sForJudge += repS
		}
		if sForJudge == s {
			return true
		}

	}

	return false
}

func repeatedSubstringPattern(s string) bool {

	if len(s) <= 1 {
		return false
	}
	for index := 0; index < len(s); index++ {
		if isRepeat(s[:index], s) {
			return true
		}
	}
	return false
}

func isRepeat(a, s string) bool {
	if len(a) == 0 {
		return false
	}
	var index = 0
	for index < len(s) {
		if strings.Index(s[index:], a) == 0 {
			index += len(a)
		} else {
			return false
		}
	}
	return true
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

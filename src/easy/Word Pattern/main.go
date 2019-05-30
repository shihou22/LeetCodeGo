package main

import "strings"

func main() {

}

func wordPattern(pattern string, str string) bool {
	s := strings.Split(str, " ")
	if len(pattern) != len(s) {
		return false
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(pattern); j++ {
			if (pattern[i] == pattern[j] && s[i] == s[j]) || (pattern[i] != pattern[j] && s[i] != s[j]) {
				continue
			} else {
				return false
			}
		}

	}
	return true
}

func wordPatternMap(pattern string, str string) bool {

	memo := make(map[rune]string)
	memoReverse := make(map[string]rune)
	stringslice := strings.Split(str, " ")
	if len(pattern) != len(stringslice) {
		return false
	}
	for index, val := range pattern {
		if tVal, ok := memo[val]; ok && tVal != stringslice[index] {
			return false
		} else if !ok {
			memo[val] = stringslice[index]
		}
		if tVal, ok := memoReverse[stringslice[index]]; ok && tVal != val {
			return false
		} else if !ok {
			memoReverse[stringslice[index]] = val
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

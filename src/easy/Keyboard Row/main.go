package main

import "strings"

func main() {

}

func findWords(words []string) []string {
	keys1 := "qwertyuiopQWERTYUIOP"
	keys2 := "asdfghjklASDFGHJKL"
	keys3 := "zxcvbnmZXCVBNM"
	// res := make([]string, 0)
	for index := 0; index < len(words); index++ {

		cnt := 0
		for _, inner := range words[index] {
			if strings.ContainsRune(keys1, inner) {
				cnt |= 1
			}
			if strings.ContainsRune(keys2, inner) {
				cnt |= 2
			}
			if strings.ContainsRune(keys3, inner) {
				cnt |= 4
			}
			if cnt != 1 && cnt != 2 && cnt != 4 {
				break
			}
		}
		if cnt != 1 && cnt != 2 && cnt != 4 {
			words = append(words[:index], words[index+1:]...)
			index--
		}
		// if cnt == 1 || cnt == 2 || cnt == 4 {
		// 	res = append(res, val)
		// }
	}
	// return res
	return words
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

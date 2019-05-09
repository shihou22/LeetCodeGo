package main

import (
	"strconv"
)

func main() {

}

func addBinaryExtra(a string, b string) string {

	var i, j, ans, carry = len(a), len(b), "", 0
	for i > 0 || j > 0 {
		num1, num2, total := 0, 0, 0
		if i > 0 {
			num1, _ = strconv.Atoi(string(a[i-1]))
			i--
		}
		if j > 0 {
			num2, _ = strconv.Atoi(string(b[j-1]))
			j--
		}
		total = num1 + num2 + carry
		switch total {
		case 0:
			ans = strconv.Itoa(0) + ans
			carry = 0
		case 1:
			ans = strconv.Itoa(1) + ans
			carry = 0
		case 2:
			ans = strconv.Itoa(0) + ans
			carry = 1
		case 3:
			ans = strconv.Itoa(1) + ans
			carry = 1
		}
	}
	if carry > 0 {
		ans = strconv.Itoa(1) + ans
	}
	return ans
}

func addBinary(a string, b string) string {
	var res string
	var carry = 0
	var max = len(a)
	if len(a) < len(b) {
		max = len(b)
	}

	for index := 0; index < max; index++ {
		var aWk = 0
		if index < len(a) {
			aWk, _ = strconv.Atoi(a[len(a)-1-index : len(a)-index])
		}
		var bWk = 0
		if index < len(b) {
			bWk, _ = strconv.Atoi(b[len(b)-1-index : len(b)-index])
		}
		var total = aWk + bWk + carry

		switch total {
		case 0:
			res = res + strconv.Itoa(0)
			carry = 0
		case 1:
			res = res + strconv.Itoa(1)
			carry = 0
		case 2:
			res = res + strconv.Itoa(0)
			carry = 1
		case 3:
			res = res + strconv.Itoa(1)
			carry = 1
		}
	}
	for index := 0; index < carry; index++ {
		res = res + strconv.Itoa(1)
	}
	return Reverse(res)
}

//Reverse string reverse
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

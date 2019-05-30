package main

import (
	"strconv"
)

func main() {

}

func addStrings(num1 string, num2 string) string {
	wkNum1 := []byte(num1)
	wkNum2 := []byte(num2)

	if len(num1) < len(num2) {
		wkNum1 = []byte(num2)
		wkNum2 = []byte(num1)
	}

	carry := 0
	for index := 0; index < len(wkNum1); index++ {

		second := 0
		if len(wkNum2)-1-index >= 0 {
			second = int(wkNum2[len(wkNum2)-1-index] - '0')
		}
		first := int(wkNum1[len(wkNum1)-1-index] - '0')
		sum := first + second + carry
		if sum > 9 {
			carry = 1
			wkNum1[len(wkNum1)-1-index] = byte(sum%10 + '0')
		} else {
			carry = 0
			wkNum1[len(wkNum1)-1-index] = byte(sum + '0')
		}
	}

	res := string(wkNum1)
	if carry > 0 {
		res = "1" + res
	}
	return res
}

func addStringsByte(num1 string, num2 string) string {
	wkNum1 := num1
	wkNum2 := num2

	if len(num1) < len(num2) {
		wkNum1 = num2
		wkNum2 = num1
	}

	res := ""
	carry := 0
	for index := range wkNum1 {
		var tmp1 int
		var tmp2 int
		tmp1 = int(wkNum1[len(wkNum1)-1-index] - '0')
		if len(wkNum2)-1-index >= 0 {
			tmp2 = int(wkNum2[len(wkNum2)-1-index] - '0')
		}
		res += strconv.Itoa((tmp1 + tmp2 + carry) % 10)
		carry = (tmp1 + tmp2 + carry) / 10
	}

	for carry != 0 {
		res += strconv.Itoa(carry % 10)
		carry = carry / 10
	}
	runes := []rune(res)
	for index := 0; index < len(runes)/2; index++ {
		runes[index], runes[len(runes)-1-index] = runes[len(runes)-1-index], runes[index]
	}

	return string(runes)
}

func addStringsRune(num1 string, num2 string) string {
	wkNum1 := []rune(num1)
	wkNum2 := []rune(num2)

	if len(num1) < len(num2) {
		wkNum1 = []rune(num2)
		wkNum2 = []rune(num1)
	}

	res := ""
	carry := 0
	for index := range wkNum1 {
		var tmp1 int
		var tmp2 int
		tmp1 = int(wkNum1[len(wkNum1)-1-index] - '0')
		if len(wkNum2)-1-index >= 0 {
			tmp2 = int(wkNum2[len(wkNum2)-1-index] - '0')
		}
		res += strconv.Itoa((tmp1 + tmp2 + carry) % 10)
		carry = (tmp1 + tmp2 + carry) / 10
	}

	for carry != 0 {
		res += strconv.Itoa(carry % 10)
		carry = carry / 10
	}
	runes := []rune(res)
	for index := 0; index < len(runes)/2; index++ {
		runes[index], runes[len(runes)-1-index] = runes[len(runes)-1-index], runes[index]
	}

	return string(runes)
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

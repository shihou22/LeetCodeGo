package main

import (
	"strconv"
)

func main() {

}
func countAndSayOther(n int) string {

	var res = "1"
	for outer := 1; outer < n; outer++ {
		var tmp = res
		res = ""
		for inner := 0; inner < len(tmp); {
			var num = tmp[inner : inner+1]
			var cnt = 0
			for k := inner; k < len(tmp); k++ {
				if num != tmp[k:k+1] {
					break
				}
				inner++
				cnt++
			}
			res += strconv.Itoa(cnt) + num
		}
	}
	return res
}

func countAndSay(n int) string {

	var wk = make([]int, 0)
	wk = append(wk, 1)
	for outer := 1; outer < n; outer++ {
		var innerWk = make([][]int, 0)
		var indexWk = 0
		for index, val := range wk {
			if index == 0 {
				var temp = make([]int, 0)
				temp = append(temp, val)
				temp = append(temp, 1)
				innerWk = append(innerWk, temp)
			} else {
				if wk[index-1] == val {
					innerWk[indexWk][1]++
				} else {
					var temp = make([]int, 0)
					temp = append(temp, val)
					temp = append(temp, 1)
					innerWk = append(innerWk, temp)
					indexWk++
				}
			}
		}
		wk = make([]int, 0)

		for _, val := range innerWk {
			wk = append(wk, val[1])
			wk = append(wk, val[0])
		}
	}

	var res string
	for _, val := range wk {
		res += strconv.Itoa(val)
	}
	return res
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

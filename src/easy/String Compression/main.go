package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func hoge(a int) {
	bts := []byte(strconv.Itoa(10))
	fmt.Println(bts)
}

func compress(chars []byte) int {

	for index := 0; index < len(chars); index++ {
		var c = chars[index]
		var cnt = 0
		for j := index + 1; j < len(chars); j++ {
			if c != chars[j] {
				break
			}
			cnt++
		}

		if cnt > 1 {
			/*
				余計なものを除去(a,a,a,a,b,b,b,b -> a,b,b,b,)
				appendで二つのslicceを一つにする
				この場合、
				　chars[:index+1]
				　chars[index+cnt:]
				この二つのslice（正確には、ひとつのsliceを二つに分割して再度結合する）
			*/
			chars = append(chars[:index+1], chars[index+cnt:]...)
			var wk = []byte(strconv.Itoa(cnt))
			for j := 0; j < len(wk); j++ {
				/*
					後ろに一つ追加
				*/
				chars = append(chars, 0)
				/*
					開始位置がindex以降の配列をindex+1に移す
				*/
				copy(chars[index+1:], chars[index:])
				/*
					空いた位置に追加
				*/
				chars[index+1] = wk[j]
				/*
					間を広げているのでindexを+1
				*/
				index++
			}

		}
	}
	return len(chars)
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

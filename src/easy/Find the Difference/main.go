package main

func main() {

}
func findTheDifference(s string, t string) byte {

	var ss [26]int
	var tt [26]int

	for index := 0; index < len(s); index++ {
		ss[int(s[index])-int('a')]++
	}
	for index := 0; index < len(t); index++ {
		tt[int(t[index])-int('a')]++
	}
	for index := 0; index < 26; index++ {
		ss[index] -= tt[index]
		if ss[index] != 0 {
			return byte('a' + index)
		}
	}
	return 'A'
}
func findTheDifferenceXor2(s string, t string) byte {

	var memo byte
	memo = 0
	for index := range t {
		if index < len(s) {
			memo ^= t[index] ^ s[index]
		} else {
			memo ^= t[index]
		}
	}
	return memo
}
func findTheDifferenceXor(s string, t string) byte {

	var memo byte
	memo = 0
	for index := range t {
		memo ^= t[index]
	}
	for index := range s {
		memo ^= s[index]
	}
	return memo
}
func findTheDifferenceMap(s string, t string) byte {

	memo := make(map[byte]int)
	for index := range s {
		memo[s[index]]++
	}
	for index := range t {
		if num, ok := memo[t[index]]; !ok || ok && num == 0 {
			return t[index]
		} else {
			memo[t[index]]--
		}
	}
	return ' '
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

package main

func main() {

}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	wk := make(map[byte]int)
	for index := 0; index < len(s); index++ {
		wk[s[index]]++
		wk[t[index]]--
	}
	for _, val := range wk {
		if val != 0 {
			return false
		}
	}
	return true
}
func isAnagramSlow(s string, t string) bool {

	sM := make(map[rune]int)
	for _, sV := range s {
		if val, ok := sM[sV]; ok {
			sM[sV] = val + 1
		} else {
			sM[sV] = 1
		}
	}
	tM := make(map[rune]int)
	for _, tV := range t {
		if val, ok := tM[tV]; ok {
			tM[tV] = val + 1
		} else {
			tM[tV] = 1
		}
	}
	for skey, sval := range sM {
		if val, ok := tM[skey]; (ok && val != sval) || (!ok) {
			return false
		}
	}
	for tkey, tval := range tM {
		if val, ok := sM[tkey]; (ok && val != tval) || (!ok) {
			return false
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

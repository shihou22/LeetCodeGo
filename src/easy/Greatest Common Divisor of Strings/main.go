package main

func main() {

}

func gcdOfStrings(str1 string, str2 string) string {
	strLen := getGCD(len(str1), len(str2))
	strBase := str2[:strLen]

	for index := 0; index < len(str1)/strLen || index < len(str2)/strLen; index += strLen {
		if index < len(str1)/strLen && strBase != str1[index:index+strLen] {
			return ""
		}
		if index < len(str2)/strLen && strBase != str2[index:index+strLen] {
			return ""
		}
	}
	return strBase
}

func getGCD(s1, s2 int) int {
	wk1 := s1
	wk2 := s2
	if s1 < s2 {
		wk1 = s2
		wk2 = s1
	}
	amari := wk1 % wk2
	for amari != 0 {
		wk1 = wk2
		wk2 = amari
		amari = wk1 % wk2
	}
	return wk2
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

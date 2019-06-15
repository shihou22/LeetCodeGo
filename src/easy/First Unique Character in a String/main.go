package main

func main() {

}

func firstUniqChar(s string) int {
	result := make([]int, 26)
	for _, val := range s {
		result[val-'a']++
	}
	for i, val := range s {
		if result[val-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqCharMap(s string) int {

	if len(s) == 1 {
		return 0
	}
	memo := make(map[rune]int)
	for _, val := range s {
		if _, ok := memo[val]; ok {
			memo[val]++
		} else {
			memo[val] = 1
		}
	}

	for i, val := range s {
		if num, ok := memo[val]; ok && num == 1 {
			return i
		}
	}
	return -1
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

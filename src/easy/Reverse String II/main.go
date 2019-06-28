package main

func main() {

}

func reverseStr(s string, k int) string {
	wk := []rune(s)

	for index := 0; index < len(wk); index += 2 * k {
		start := index
		end := index + k - 1
		if end > len(wk)-1 {
			end = len(wk) - 1
		}

		for start < end {
			wk[start], wk[end] = wk[end], wk[start]
			start++
			end--
		}

	}

	return string(wk)
}

func reverseStrOld(s string, k int) string {
	wk := []rune(s)
	/*
		abcdefg
		↑  ↑  ↑
		 ↑  ↑  ↑
	*/
	for index := 0; index < len(wk); index += 2 * k {
		prev := index
		target := index + k - 1
		if target > len(wk)-1 {
			target = len(wk) - 1
		}
		for prev < target {
			wk[prev], wk[target] = wk[target], wk[prev]
			prev++
			target--
		}
	}
	return string(wk)
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

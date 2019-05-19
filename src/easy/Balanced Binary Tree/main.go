package main

func main() {

}

func isBalanced(root *TreeNode) bool {
	var _, ok = isBalancedTree(root)
	return ok
}

func isBalancedTree(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	var depthL, okL = isBalancedTree(root.Left)
	var depthR, okR = isBalancedTree(root.Right)
	if !okL || !okR {
		return 0, false
	}
	if abs(depthL-depthR) <= 1 {
		return max(depthL, depthR) + 1, true
	}
	return 0, false
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
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

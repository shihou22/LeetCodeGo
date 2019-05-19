package main

func main() {

}
func sortedArrayToBST(nums []int) *TreeNode {
	return makeNode(nums, 0, len(nums)-1)
}

/*
0 1 2 3 4 5 6 7 8 9 10
0 1 2 3 4   6 7 8 9 10
0 1   3 4   6 7   9 10
0     3     6     9
*/
func makeNode(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return &TreeNode{nums[start], nil, nil}
	}
	var mid = start + ((end - start) >> 1)
	// var mid int
	// if end-start < 2 {
	// 	mid = (start+end)/2 + 1
	// } else {
	// 	mid = (start + end) / 2
	// }

	var node = &TreeNode{nums[mid], nil, nil}
	var left = makeNode(nums, start, mid-1)
	var right = makeNode(nums, mid+1, end)

	node.Left = left
	node.Right = right
	return node
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

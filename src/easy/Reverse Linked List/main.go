package main

func main() {

}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var reverseListRecursive func(head *ListNode, root *ListNode) *ListNode
	reverseListRecursive = func(head *ListNode, root *ListNode) *ListNode {
		if head == nil {
			return root
		}
		var list = &ListNode{head.Val, root}
		return reverseListRecursive(head.Next, list)
	}

	return reverseListRecursive(head.Next, &ListNode{head.Val, nil})
}

func reverseListIterable(head *ListNode) *ListNode {
	var list *ListNode

	for head != nil {
		list = &ListNode{Val: head.Val, Next: list}
		head = head.Next
	}

	return list
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

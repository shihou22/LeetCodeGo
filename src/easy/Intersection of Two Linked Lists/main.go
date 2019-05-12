package main

func main() {

}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	wkA := headA
	wkB := headB
	if wkA == nil || wkB == nil {
		return nil
	}
	var aLen = 0
	for wkA != nil {
		wkA = wkA.Next
		aLen++
	}
	var bLen = 0
	for wkB != nil {
		wkB = wkB.Next
		bLen++
	}

	var aheadMax = 0
	var ahead1 *ListNode
	var ahead2 *ListNode
	if aLen > bLen {
		aheadMax = aLen - bLen
		ahead1 = headA
		ahead2 = headB
	} else {
		aheadMax = bLen - aLen
		ahead1 = headB
		ahead2 = headA
	}
	for index := 0; index < aheadMax; index++ {
		ahead1 = ahead1.Next
	}

	for ahead1 != ahead2 {
		ahead1 = ahead1.Next
		ahead2 = ahead2.Next
	}
	return ahead1
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

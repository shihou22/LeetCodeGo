package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var res = &ListNode{}
	var resTmp = res
	var carry = 0
	for l1 != nil || l2 != nil {
		var l1V = 0
		var l2V = 0
		if l1 != nil {
			l1V = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2V = l2.Val
			l2 = l2.Next
		}
		var tmp = l1V + l2V + carry
		if tmp >= 10 {
			carry = 1
			tmp = tmp % 10
		} else {
			carry = 0
		}
		var wk = ListNode{}
		wk.Val = tmp
		resTmp.Next = &wk
		resTmp = resTmp.Next
	}
	if carry > 0 {
		resTmp.Next = &ListNode{Val: carry, Next: nil}
	}
	return res.Next
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

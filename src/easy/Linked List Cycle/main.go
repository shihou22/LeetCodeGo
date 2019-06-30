package main

func main() {

}

func hasCycleOther(head *ListNode) bool {

	var base = head
	var next = head
	for base != nil {
		if next == nil || next.Next == nil {
			return false
		}
		base = base.Next
		next = next.Next.Next
		if base == next {
			return true
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}
	memo := make(map[*ListNode]int)
	for head != nil {
		if _, ok := memo[head]; !ok {
			memo[head]++
		} else {
			return true
		}
		head = head.Next
	}
	return false
}

func hasCycleOld(head *ListNode) bool {
	var wk = make(map[*ListNode]int)
	for head != nil {
		var _, ok = wk[head]
		if ok {
			return true
		} else {
			wk[head] = 1
		}
		head = head.Next
	}
	return false
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

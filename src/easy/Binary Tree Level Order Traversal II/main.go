package main

func main() {

}

func levelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	var box = make(map[int][]int)
	levelOrderBottomRecursive(root, box, 0)

	var res = make([][]int, 0)
	var max = -1
	for index := 0; index < len(box); index++ {
		var tmp = make([]int, 0)
		res = append(res, tmp)
		max++
	}

	for key, val := range box {
		res[max-key] = val
	}

	return res
}

func levelOrderBottomRecursive(root *TreeNode, box map[int][]int, current int) {
	if root == nil {
		return
	}

	var res, ok = box[current]
	if ok {
		res = append(res, root.Val)
		box[current] = res
	} else {
		var tmp = make([]int, 0)
		tmp = append(tmp, root.Val)
		box[current] = tmp
	}
	if root.Left != nil {
		levelOrderBottomRecursive(root.Left, box, current+1)
	}
	if root.Right != nil {
		levelOrderBottomRecursive(root.Right, box, current+1)
	}
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

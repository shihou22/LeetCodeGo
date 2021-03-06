package main

func main() {

}

func pathSum(root *TreeNode, sum int) int {

	if root == nil {
		return 0
	}
	memo := make([]int, 0)
	return pathSumHelper(root, sum, memo)
}

func pathSumHelper(root *TreeNode, sum int, memo []int) int {

	if root == nil {
		return 0
	}
	memo = append(memo, root.Val)
	val0 := 0
	total := 0
	for index := len(memo) - 1; index >= 0; index-- {
		total += memo[index]
		if total == sum {
			val0++
			/*
				doesn't brek because in some case val is 0 or minus
			*/
		}
	}

	val1 := pathSumHelper(root.Left, sum, memo)
	val2 := pathSumHelper(root.Right, sum, memo)
	return val0 + val1 + val2
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

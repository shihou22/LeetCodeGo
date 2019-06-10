package main

func main() {

}

var result []int

func findModeMLE(root *TreeNode) []int {
	result = make([]int, 0)
	getModeCnt(root)
	max := 0
	for _, val := range result {
		if max < val {
			max = val
		}
	}
	cnt := 1
	for index := 0; index < len(result); index++ {
		if max != result[index] {
			result = result[index+1:]
			index--
		} else {
			result[index] = cnt
		}
		cnt++
	}

	return result
}

func getModeCnt(root *TreeNode) {

	if root == nil {
		return
	}
	if root.Val > len(result) {
		for index := 0; index < root.Val-len(result); index++ {
			result = append(result, 0)
		}
	}
	result[root.Val-1]++
	getModeCnt(root.Left)
	getModeCnt(root.Right)
}

func findModeMap(root *TreeNode) []int {
	memo := make(map[int]int)
	max := getModeCntMap(root, memo)
	res := make([]int, 0)
	for key, val := range memo {
		if max == val {
			res = append(res, key)
		}
	}
	return res
}

func getModeCntMap(root *TreeNode, memo map[int]int) int {

	if root == nil {
		return 0
	}
	if _, ok := memo[root.Val]; ok {
		memo[root.Val]++
	} else {
		memo[root.Val] = 1
	}

	return max(memo[root.Val], getModeCntMap(root.Left, memo), getModeCntMap(root.Right, memo))

}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else if c >= a && c >= b {
		return c
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

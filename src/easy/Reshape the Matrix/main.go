package main

func main() {

}
func matrixReshape(nums [][]int, r int, c int) [][]int {

	if nums == nil || len(nums) == 0 || len(nums)*len(nums[0]) > r*c {
		return nums
	}

	if len(nums)*len(nums[0]) <= r*c && len(nums) <= r && len(nums[0]) <= c {
		c = len(nums[0])
	}

	memo := make([]int, len(nums)*len(nums[0]))
	d := 0
	for _, val1 := range nums {
		for _, val2 := range val1 {
			memo[d] = val2
			d++
		}
	}

	res := make([][]int, 0)
	d = 0
	for outer := 0; outer < r && d < len(memo); outer++ {
		tmp := make([]int, 0)
		res = append(res, tmp)
		for inner := 0; inner < c; inner++ {
			if d < len(memo) {
				res[outer] = append(res[outer], memo[d])
				d++
			}
		}
	}
	return res

}

func matrixReshapeComlex(nums [][]int, r int, c int) [][]int {

	if nums == nil || len(nums)*len(nums[0]) > r*c {
		return nums
	}
	if len(nums)*len(nums[0]) <= r*c && len(nums) <= r && len(nums[0]) <= c {
		c = len(nums[0])
	}
	res := make([][]int, 0)
	wkR := 0
	wkC := 0
	for _, rVal := range nums {
		for _, cVal := range rVal {
			if wkR != len(res)-1 {
				tmp := make([]int, 0)
				res = append(res, tmp)
			}
			res[wkR] = append(res[wkR], cVal)
			wkC++
			if wkC == c {
				wkC = 0
				wkR++
			}
		}
	}
	return res
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

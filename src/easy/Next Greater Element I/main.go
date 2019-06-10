package main

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	memo := make(map[int]int)
	s := make([]int, 0)
	for _, num := range nums2 {
		for len(s) > 0 && s[len(s)-1] < num {
			key := s[len(s)-1]
			s = append(s[:len(s)-1])
			memo[key] = num
		}
		s = append(s, num)
	}

	res := make([]int, 0)
	for _, val := range nums1 {
		if num, ok := memo[val]; ok {
			res = append(res, num)
		} else {
			res = append(res, -1)
		}
	}
	return res
}

func nextGreaterElementOn2(nums1 []int, nums2 []int) []int {

	res := make([]int, 0)
	for _, base := range nums1 {
		containsI := -1
		resVal := -1
		for cI, cVal := range nums2 {
			if cVal == base {
				containsI = cI
			}
			if containsI != -1 && cVal > base {
				resVal = cVal
				break
			}
		}
		res = append(res, resVal)

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

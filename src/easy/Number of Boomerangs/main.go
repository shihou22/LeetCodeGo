package main

func main() {

}

func numberOfBoomerangs(points [][]int) int {

	res := 0
	memo := make(map[int]int)
	for outerI, outerVal := range points {
		for innerI, innerVal := range points {
			if innerI == outerI {
				continue
			}
			memo[calcDistance(outerVal, innerVal)]++
		}
		for _, val := range memo {
			/*
				中心を固定してnP2を行う
				等距離に1個しかない場合は0になるので無視
			*/
			res += val * (val - 1)
		}
		memo = make(map[int]int)
	}

	return res
}

func calcDistance(a, b []int) int {
	xVal := a[0] - b[0]
	yVal := a[1] - b[1]
	return xVal*xVal + yVal*yVal
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

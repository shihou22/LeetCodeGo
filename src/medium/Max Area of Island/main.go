package main

func main() {

}

/*
TreeNode
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for oI, oV := range grid {
		for iI, iV := range oV {
			if iV == 1 {
				cnt := countArea(grid, oI, iI)
				res = Max(res, cnt)
			}
		}
	}
	return res
}

func countArea(grid [][]int, y int, x int) int {

	d := []int{0, 1, 0, -1}
	res := 0
	grid[y][x] = 2
	res++
	for index := 0; index < len(d); index++ {
		wkY := d[index] + y
		wkX := d[index^1] + x
		if wkY < 0 || wkY >= len(grid) || wkX < 0 || wkX >= len(grid[0]) || grid[wkY][wkX] != 1 {
			continue
		}
		grid[wkY][wkX] = 2
		res += countArea(grid, wkY, wkX)
	}
	return res
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

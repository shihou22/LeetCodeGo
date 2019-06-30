package main

func main() {

}

func numIslands(grid [][]byte) int {

	visited := make([][]bool, 0)
	for _, val := range grid {
		tmp := make([]bool, len(val))
		visited = append(visited, tmp)
	}
	res := 0
	for oI := 0; oI < len(grid); oI++ {
		for iI := 0; iI < len(grid[oI]); iI++ {
			if !visited[oI][iI] && grid[oI][iI] != '0' {
				visited[oI][iI] = true
				recursive(grid, visited, oI, iI)
				res++
			}
		}
	}
	return res
}

func recursive(grid [][]byte, visited [][]bool, oI int, iI int) {
	d := []int{0, 1, 0, -1}
	for dI := 0; dI < len(d); dI++ {
		wkY := d[dI^1] + oI
		wkX := d[dI] + iI
		if wkX < 0 || wkX >= len(grid[0]) || wkY < 0 || wkY >= len(grid) || visited[wkY][wkX] || grid[wkY][wkX] == '0' {
			continue
		}
		visited[wkY][wkX] = true
		recursive(grid, visited, wkY, wkX)
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

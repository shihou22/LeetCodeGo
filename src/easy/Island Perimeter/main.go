package main

func main() {

}

type Dimension struct {
	x int
	y int
}

func getDimension() []Dimension {
	dimension := make([]Dimension, 4)
	tmp := Dimension{}
	tmp.x = -1
	tmp.y = 0
	dimension[0] = tmp
	tmp = Dimension{}
	tmp.x = 1
	tmp.y = 0
	dimension[1] = tmp
	tmp = Dimension{}
	tmp.x = 0
	tmp.y = 1
	dimension[2] = tmp
	tmp = Dimension{}
	tmp.x = 0
	tmp.y = -1
	dimension[3] = tmp
	return dimension
}

func islandPerimeter(grid [][]int) int {
	if grid == nil {
		return 0
	}
	x := len(grid[0])
	y := len(grid)
	// dimension := getDimension()
	res := 0

	dX := []int{-1, 1, 0, 0}
	dY := []int{0, 0, -1, 1}

	for gridY, gridVal := range grid {

		for gridX, inVal := range gridVal {
			if inVal == 0 {
				continue
			}
			// for _, inDVal := range dimension {
			for index := 0; index < 4; index++ {
				tmpX := gridX + dX[index]
				tmpY := gridY + dY[index]
				// tmpX := gridX + inDVal.x
				// tmpY := gridY + inDVal.y
				if tmpX >= 0 && tmpX < x && tmpY >= 0 && tmpY < y && grid[tmpY][tmpX] == 0 {
					res++
				} else if tmpX < 0 || tmpX >= x || tmpY < 0 || tmpY >= y {
					res++
				}
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

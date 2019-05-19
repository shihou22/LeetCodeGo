package main

func main() {

}

func getRow(rowIndex int) []int {

	var res = make([]int, 0)
	res = append(res, 1)
	for index := 0; index < rowIndex; index++ {
		for inner := index + 1; inner > 0; inner-- {

			if inner == index+1 {
				res = append(res, 1)
			} else {
				res[inner] = res[inner] + res[inner-1]
			}
		}
	}
	return res
}

func getRowFirst(rowIndex int) []int {

	var res = make([][]int, 0)

	for index := 0; index <= rowIndex; index++ {
		var temp = make([]int, 0)
		temp = append(temp, 1)
		res = append(res, temp)
	}
	for index := 1; index <= rowIndex; index++ {
		var prev = res[index-1]
		var temp = res[index]
		for inner := 1; inner < index+1; inner++ {
			if inner == index {
				temp = append(temp, 1)
			} else {
				temp = append(temp, prev[inner]+prev[inner-1])
			}
		}
		res[index] = temp
	}
	return res[rowIndex]
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

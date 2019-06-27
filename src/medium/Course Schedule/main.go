package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	neighborList := make(map[int][]int)
	for _, Val := range prerequisites {
		var tmp []int
		if vertic, ok := neighborList[Val[1]]; !ok {
			tmp = make([]int, 0)
		} else {
			tmp = vertic
		}
		tmp = append(tmp, Val[0])
		neighborList[Val[1]] = tmp
	}

	vertics := make([]int, numCourses)
	for _, Val := range neighborList {
		for _, vertic := range Val {
			vertics[vertic]++
		}
	}

	queue := make([]int, 0)

	for Index, verticNum := range vertics {
		if verticNum == 0 {
			queue = append(queue, Index)
		}
	}

	if len(queue) == 0 {
		return false
	}
	cnt := 0

	for len(queue) != 0 {

		pos := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for _, verticNum := range neighborList[pos] {
			vertics[verticNum]--
			if vertics[verticNum] == 0 {
				queue = append(queue, verticNum)
			}
		}
		cnt++
		if cnt > numCourses {
			return false
		}
	}
	return cnt == numCourses
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

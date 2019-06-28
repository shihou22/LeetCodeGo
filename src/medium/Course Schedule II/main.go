package main

func main() {

}

func findOrder(numCourses int, prerequisites [][]int) []int {

	neighbors := make(map[int][]int)
	for _, Val := range prerequisites {
		var tmp []int
		if _, ok := neighbors[Val[1]]; ok {
			tmp = neighbors[Val[1]]
		} else {
			tmp = make([]int, 0)
		}
		tmp = append(tmp, Val[0])
		neighbors[Val[1]] = tmp
	}

	degrees := make([]int, numCourses)
	for _, verticles := range neighbors {
		for _, verticle := range verticles {
			degrees[verticle]++
		}
	}

	queue := make([]int, 0)
	for Index, degree := range degrees {
		if degree == 0 {
			queue = append(queue, Index)
		}
	}

	res := make([]int, 0)
	cnt := 0
	for len(queue) != 0 {

		pos := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		list := neighbors[pos]
		for _, Val := range list {
			degrees[Val]--
			if degrees[Val] == 0 {
				queue = append(queue, Val)
			}
		}
		cnt++
		res = append(res, pos)
		if cnt > numCourses {
			return []int{}
		}
	}
	if cnt == numCourses {
		return res
	} else {
		return []int{}
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

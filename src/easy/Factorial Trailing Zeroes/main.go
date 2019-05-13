package main

func main() {

}

func trailingZeroes(n int) int {
	var five = 0
	for {
		if n/5 < 1 {
			return five
		} else {
			five += n / 5
			n /= 5
		}
	}
}

func trailingZeroesBrute(n int) int {

	var wk = make(map[int]int)

	for index := 1; index <= n; index++ {
		var target = index
		for inner := 2; inner <= n; {
			if target%inner == 0 {
				wk[inner]++
				target /= inner
			} else {
				inner++
			}
		}
	}
	var val, ok = wk[5]
	if ok {
		return val
	} else {
		return 0
	}
}

func trailingZeroesOther(n int) int {

	var res int64
	res = 1
	for index := 1; index <= n; index++ {
		res *= int64(index)
	}

	var sum = 0
	remainder := int64(0)
	for remainder == 0 {
		remainder = res % 10
		if remainder == 0 {
			sum++
		}
		res /= 10
	}
	return sum
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

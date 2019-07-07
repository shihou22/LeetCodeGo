package main

func main() {

}

/*
https://kazu-yamamoto.hatenablog.jp/entry/20090223/1235372875
http://augusuto04.hatenablog.com/entry/2015/05/02/183451
http://ak-blog.hatenablog.jp/entry/20130426/1366985322
*/
func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	// V3
	var ret float64
	ret = 1
	for 0 < n {
		if (n % 2) == 0 {
			x *= x
			n >>= 1
		} else {
			ret *= x
			n--
		}
	}
	return ret

	//V2
	// if n%2 == 0 {
	// 	return myPow(x*x, n/2)
	// }
	// return x * myPow(x, (n-1))

	// V1
	// res := myPow(x, n/2)
	// if n%2 == 0 {
	// 	return res * res
	// }
	// return res * res * x

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

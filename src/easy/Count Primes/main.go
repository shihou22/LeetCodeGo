package main

import "math"

func main() {

}
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	var res = 0
	var memo = make([]bool, n+1)
	for index := 2; index < n; index++ {
		if !memo[index] {
			memo[index] = true
			res++
			for inner := 2; inner*index < n; inner++ {
				memo[inner*index] = true
			}
		}
	}
	return res
}
func countPrimesUseMap(n int) int {
	if n <= 2 {
		return 0
	}
	var res = 0
	var memo = make(map[int]bool)
	for index := 2; index < n; index++ {
		if _, ok := memo[index]; ok {
			continue
		} else {
			memo[index] = true
			res++
		}
		for inner := 2; inner*index < n; inner++ {
			memo[inner*index] = true
		}
	}
	return res
}

func countPrimesSlow(n int) int {
	if n <= 2 {
		return 0
	}
	var res = 1
	for index := 3; index < n; index += 2 {
		var max = int(math.Sqrt(float64(index)))
		if index%2 == 0 {
			continue
		}
		var isPrime = true
		for inner := 3; inner <= max; inner++ {
			if index%inner == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			res++
		}
	}
	return res
}

func countPrimesNormal(n int) int {

	if n <= 2 {
		return 0
	}
	var res = 1
	for index := 3; index < n; index += 2 {
		var isPrime = true
		for inner := 2; inner < index; inner++ {
			if index%inner == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			res++
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

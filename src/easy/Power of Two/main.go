package main

func main() {

}
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%2 == 0 {
		return isPowerOfTwo(n / 2)
	}
	return false
}
func isPowerOfTwoOther(n int) bool {

	if n == 1 {
		return true
	}
	var res = 1
	for index := 1; index <= 32; index++ {
		if n == res {
			return true
		}
		res <<= 1
		if n < res {
			return false
		}
	}
	return false
	// if n == 1 {
	// 	return true
	// }

	// //this argument is Integer. so Number of digits are 32.
	// for index := 1; index <= 32; index++ {
	// 	if n == (1 << uint(index)) {
	// 		return true
	// 	}
	// }
	// return false
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

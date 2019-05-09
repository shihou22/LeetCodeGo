package main

func main() {

}
func plusOne(digits []int) []int {

	var copyA = make([]int, len(digits)+1)

	var plusOne = 1
	for index := 0; index < len(digits); index++ {
		var carry = copyA[len(copyA)-1-index] + digits[len(digits)-1-index] + plusOne
		if carry >= 10 {
			copyA[len(copyA)-1-index] = carry % 10
			copyA[len(copyA)-2-index]++
		} else {
			copyA[len(copyA)-1-index] = carry
		}
		carry = 0
		plusOne = 0
	}
	if copyA[0] == 0 {
		return append(copyA[1:len(copyA)])
	} else {
		return copyA
	}

	// var wk int64
	// wk = 0
	// for index := 0; index < len(digits); index++ {
	// 	wk *= 10
	// 	wk += int64(digits[index])
	// }
	// wk++
	// var checkD = wk
	// var dig int64
	// dig = 0
	// for checkD != 0 {
	// 	checkD /= 10
	// 	dig++
	// }
	// checkD = wk
	// var res = make([]int, dig)
	// for index := int64(0); index < dig; index++ {
	// 	res[dig-1-index] = int(checkD % 10)
	// 	checkD /= 10
	// }
	// return res
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

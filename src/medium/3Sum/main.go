package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {

}

func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue //To prevent the repeat
		}

		Left := -nums[i]
		Middle := i + 1
		Right := len(nums) - 1

		for Middle < Right {
			sum := nums[Middle] + nums[Right]
			if sum == Left {
				results = append(results, []int{nums[i], nums[Middle], nums[Right]})
				Middle++
				for Middle < Right && nums[Middle] == nums[Middle-1] {
					Middle++
				}
				for Middle < Right && Right < len(nums)-1 && nums[Right] == nums[Right+1] {
					Right--
				}
			} else if sum > Left {
				Right--
			} else if sum < Left {
				Middle++
			}
		}

	}
	return results
}

func threeSumOld(nums []int) [][]int {
	memo := make(map[int]int)
	key := make(map[string]int)
	for _, val := range nums {
		memo[val]++
	}
	res := make([][]int, 0)
	for index1 := 0; index1 < len(nums); index1++ {
		num1 := nums[index1]
		memo[num1]--
		for index2 := index1 + 1; index2 < len(nums); index2++ {
			if index1 == index2 {
				continue
			}
			num2 := nums[index2]
			memo[num2]--
			num3 := 0 - num1 - num2
			if _, ok := memo[num3]; ok {
				memo[num3]--
				if memo[num1] >= 0 && memo[num2] >= 0 && memo[num3] >= 0 {
					tmpK := strconv.Itoa(Min(num1, num2, num3)) + ":" + strconv.Itoa(Middle(num1, num2, num3)) + ":" + strconv.Itoa(Max(num1, num2, num3))
					key[tmpK]++
				}
				memo[num3]++
			}
			memo[num2]++
		}
	}

	for elKey := range key {
		numsA := strings.Split(elKey, ":")
		tmp := make([]int, 0)
		num1, _ := strconv.Atoi(numsA[0])
		num2, _ := strconv.Atoi(numsA[1])
		num3, _ := strconv.Atoi(numsA[2])
		tmp = append(tmp, num1, num2, num3)
		res = append(res, tmp)
	}
	return res
}

func Max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func Min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func Middle(a, b, c int) int {
	if (a >= b && b >= c) || (a <= b && b <= c) {
		return b
	}
	if (b >= a && a >= c) || (b <= a && a <= c) {
		return a
	}
	return c
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

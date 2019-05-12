package main

func main() {

}

/*
defince struct
*/
type MinStack struct {
	Stack []int
}

func Constructor() MinStack {
	var body MinStack
	return body
}

func (this *MinStack) Push(x int) {
	// (*this).Stack = append((*this).Stack, x)
	// var tmp *MinStack
	// tmp = this
	// tmp.Stack = append(tmp.Stack, x)
	this.Stack = append(this.Stack, x)
}

func (this *MinStack) Pop() {
	// (*this).Stack = (*this).Stack[:len((*this).Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	min := this.Stack[0]
	for i := 0; i < len(this.Stack); i++ {
		if min > this.Stack[i] {
			min = this.Stack[i]
		}
	}
	return min
}

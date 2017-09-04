type MyStack struct {
    stack []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    stack := new(MyStack)
    return *stack
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
     this.stack = append(this.stack, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    index := len(this.stack)
    if index==0 {
        return 0
    }
    res := this.stack[index-1]
    this.stack = this.stack[:index-1]
    return res
}


/** Get the top element. */
func (this *MyStack) Top() int {
    index := len(this.stack)
    if index==0 {
        return 0
    } else {
        return this.stack[index-1]
    }
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    index := len(this.stack)
    if index==0 {
        return true
    } else {
        return false
    }
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

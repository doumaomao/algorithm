type MyQueue struct {
    queue []int
    
}
/** Initialize your data structure here. */
func Constructor() MyQueue {
    queue := new(MyQueue)
    return *queue
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.queue = append(this.queue, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    index := len(this.queue)
    if index == 0 {
        return 0
    }
    res := this.queue[0]
    if index == 1 {
        this.queue = this.queue[:index-1]
    } else {
        this.queue = this.queue[1:]
    }
    return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
       index := len(this.queue)
    if index==0 {
        return 0
    } else {
        return this.queue[0]
    }
}
/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    index := len(this.queue)
    if index==0 {
        return true
    } else {
        return false
    }
}
/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

package easy

/*
Title: 用队列实现栈
Difficulty: easy
Tags: 栈, 设计, 队列
Link: https://leetcode.cn/problems/implement-stack-using-queues/

Description:
请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

实现 MyStack 类：

	void push(int x) 将元素 x 压入栈顶。
	int pop() 移除并返回栈顶元素。
	int top() 返回栈顶元素。
	boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。

注意：

	你只能使用队列的标准操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
	你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。

示例：
输入：
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
输出：
[null, null, null, 2, 2, false]

解释：
MyStack myStack = new MyStack();
myStack.push(1);
myStack.push(2);
myStack.top(); // 返回 2
myStack.pop(); // 返回 2
myStack.empty(); // 返回 False

提示：
	1 <= x <= 9
	最多调用100 次 push、pop、top 和 empty
	每次调用 pop 和 top 都保证栈不为空

进阶：你能否仅用一个队列来实现栈。
*/

type _225_MyStack struct {
	q1, q2 []int
}

func _225_Constructor() _225_MyStack {
	return _225_MyStack{}
}

func (this *_225_MyStack) Push(x int) {
	this.q2 = append(this.q2, x)
	for len(this.q1) > 0 {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
	}
	this.q1, this.q2 = this.q2, this.q1
}

func (this *_225_MyStack) Pop() int {
	e := this.q1[0]
	this.q1 = this.q1[1:]
	return e
}

func (this *_225_MyStack) Top() int {
	return this.q1[0]
}

func (this *_225_MyStack) Empty() bool {
	return len(this.q1) == 0
}

type _225_MyStack2 struct {
	q []int
}

func _225_Constructor2() _225_MyStack2 {
	return _225_MyStack2{}
}

func (this *_225_MyStack2) Push(x int) {
	n := len(this.q)
	this.q = append(this.q, x)
	for ; n > 0; n-- {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
	}

}

func (this *_225_MyStack2) Pop() int {
	e := this.q[0]
	this.q = this.q[1:]
	return e
}

func (this *_225_MyStack2) Top() int {
	return this.q[0]
}

func (this *_225_MyStack2) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

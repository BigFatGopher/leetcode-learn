type MinStack struct {
    //普通栈（后进先出）
    stackData []int
    //辅助栈，维护一个最小元素在data栈相对顺序的栈
    stackMin []int
    //记录当前两个栈的大小，减少遍历
    lengthData, lengthMin int
}
/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stackData : []int{},
        //加一个int64最大值，可以减少对min栈为空的判断
        stackMin  : []int{math.MaxInt64},
        lengthMin : 1,
    }
}
func (this *MinStack) Push(x int)  {
    this.stackData = append(this.stackData, x)
    this.lengthData++
    //最小元素大于x时无需加入
    if this.stackMin[this.lengthMin-1] >= x{
        this.stackMin = append(this.stackMin, x)
        this.lengthMin++
    }    
}
func (this *MinStack) Pop()  {
    x := this.stackData[this.lengthData-1]
    //出栈
    this.stackData = this.stackData[:this.lengthData-1]
    this.lengthData--
    //若data栈出栈的元素是最小元素，那么min栈也需要进行弹出操作
    if this.stackMin[this.lengthMin-1] == x {
        this.stackMin = this.stackMin[:this.lengthMin-1]
        this.lengthMin--
    }
}
func (this *MinStack) Top() int {
    return this.stackData[this.lengthData-1]
}
func (this *MinStack) GetMin() int {
    return this.stackMin[this.lengthMin-1]
}
实现方式二
压入数据规则
假设当前数据为newNum，先将其压入stackData。然后判断stackMin是否为空。
如果为空，则newNum也压入stackMin；如果不为空，则比较newNum和stackMin的栈顶元素中哪一个更小：如果newNum更小或两者相等，则newNum也压入stackMin；如果stackMin中栈顶元素小，则把stackMin的栈顶元素重复压入stackMin，即在栈顶元素上再压入一个栈顶元素。
举例：依次压入3、4、5、1、2、1的过程中，stackData和stackMin的变化如下图所示。


弹出数据规则
在stackData中弹出数据，弹出的数据记为value；弹出stackMin中的栈顶；返回value。
很明显可以看出，压入与弹出规则是对应的

查询当前栈中的最小值操作
由上文的压入数据规则和弹出数据规则可知，stackMin始终记录着stackData中的最小值，所以stackMin的栈顶元素始终是当前stackData中的最小值。


type MinStack struct {
    //保存当前栈的元素
    stackData []int 
    //保存每一步的最小值
    stackMin []int
    //记录当前两个栈的大小，减少遍历
    lengthData, lengthMin int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    //给stackMin一个初始元素，是为了减少判断stackMin为空的情况
    return MinStack{stackMin: []int{math.MaxInt64}, lengthMin: 1}
}


func (this *MinStack) Push(x int)  {
    this.stackData = append(this.stackData, x)
    //压入栈，将栈顶元素和新元素比较，压入最小的元素
    this.stackMin = append(this.stackMin, min(x, this.stackMin[this.lengthMin-1]))
    this.lengthData++
    this.lengthMin++
}


func (this *MinStack) Pop()  {
    this.stackData = this.stackData[:this.lengthData-1]
    this.stackMin = this.stackMin[:this.lengthMin-1]    
    this.lengthMin--; this.lengthData--
}


func (this *MinStack) Top() int {
    return this.stackData[this.lengthData-1]
}


func (this *MinStack) GetMin() int {
    return this.stackMin[this.lengthMin-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
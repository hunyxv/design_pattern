## 策略模式
策略模式(Strategy pattern)鼓励使用多种算法来解决一个问题，其杀手级特性是能够在运行时透明地切换算法(客户端代码对变化无感知)。因此，如果你有两种算法，并且知道其中一种对少量输入效果更好，另一种对大量输入效果更好，则可以使用策略模式在运行时基于输入数据决定使用哪种算法。


官方`sort`包就是使用了策略模式，sort 包含了插入排序、堆排序、快速排序和归并排序，对使用者来说不用关心 sort 使用了那种排序方法。
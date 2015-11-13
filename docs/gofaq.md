GO语言学习笔记 - FAQ

Q: go的range在什么时候循环结束？
A: 只要不是关闭就不会结果吧。关闭后for range结束。


Q: go的channel完成控制为什么不能直接关闭？如果确实不能关闭，只能使用done channel方式吗，太费劲了吧？
A: deferchan(ch)?

Q: go的import包为什么使用字符串？
A: 总的来说，直接对应到文件系统中的$GOPATH/src下的路径。
   那么这里有可能出现跨平台的问题吧。

Q: go的抽像能力，Ducking Typing如何？
A: 

Q: 为什么不用channel的close作为完成标识？
A: 有可能是异常而close了，接收方无法知道的原因吗？
读取一个关闭的channel，返回值为channel类型的默认值，如chan string，则返回一个空字符串。
无法确定channel是否完成。
close两次channel会panic。

Q: 只读或者只写的channel有什么用？
A: 






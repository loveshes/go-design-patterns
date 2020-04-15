# Golang 实现23种设计模式

> 之前也稍微学过设计模式，现在尝试用Go来实现经典的23种设计模式，也算是一种复习了。
>
> 参考资料：[大话设计模式](https://book.douban.com/subject/2334288/) | [图解设计模式](https://book.douban.com/subject/26933281/) | [菜鸟教程—设计模式](https://www.runoob.com/design-pattern/design-pattern-tutorial.html)



## 设计模式

### 0. 简单工厂模式

[CODE](pattern/factory-pattern/factory/factory.go) | [详解](pattern/factory-pattern/README.md)

通过传入参数给工厂，让工厂来生成我们想要的实际对象，进而对对象进行一系列操作。



### 1. 策略模式

[CODE](pattern/strategy-pattern/strategy/strategy.go) | [详解](pattern/strategy-pattern/README.md)

定义了策略家族，分别封装起来，让它们之间可以互相替换，此模式让策略的变化，不会影响到使用策略的客，其实就是多态。也可与简单工厂模式结合，传入参数给工厂，由工厂来生成具体的策略。



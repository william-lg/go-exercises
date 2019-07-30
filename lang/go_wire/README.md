## 关于Go Wire使用的几点思考
### Go Wire在使用中的一些特性
1. wire.NewSet(SetA,SetB,SetC,...)会作为一个整体对调用者提供，但是要注意的是，整体的每个元素都会被参与统计。
   因此，当多个Set中的元素有相同的时候，wire会统计多次
2. wire.NewSet(funcA,funB,funcC,...)，func函数会生成新的元素供外界调用，只要其需要的各种依赖都能被提供
3. wire.Bind(interface, interfaceImpl)会生成一个新的interface元素
4. wire不允许有对同一元素的多次依赖，wire认为，如果多个调用者都需要同一元素的依赖，则其中一个对其进行了更改，
   会影响到其他调用者的使用
5. wire不允许一个类型有多个提供者，比如interface元素有两个实现者同时提供，这是不允许的

### Go Wire在使用过程中遇到的一些问题以及解决办法
使用背景: 分层架构repo->manager->service->facade | <-options(third_party<-config)
1. 同一元素多次依赖，常见于特性1中表述的使用方法，原因也已经说明  
`解决方法:`  
A. 对于repo manager third_party config层提供统一的注入全集，这样在上层调用中只需使用
wire.NewSet(funcA,funB,funcC,allRepoSet,allManagerSet...)即可，把依赖全部提供，上层
需要的只是将各种依赖合成的新的高层元素。  
B. 同时在某一层级只需注入下一层级的全部依赖，本层可以自定义Func依赖，解决了同层级的相互引用问题
2. 同一元素多个提供者，常见于接口的实现注入中  
`解决办法:`  
目前Demo里遵循的原则是在struct定义中可以定义结构体变量为interface，但是在NewFunc中需要
指定interface的实现者，而不是直接注入interface类型。这样做的原因是目前发现了interface的
嵌套等用法，有可能会存在一个interface存在多个实现者并提供。

### 说明:
以上仅为个人在使用过程中的体会，存在猜测成分。

### 常用链接:
[Google Wire Tutorials](https://github.com/google/wire/blob/master/_tutorial/README.md)  
[Google Wire User Guides](https://github.com/google/wire/blob/master/docs/guide.md)  
[Google Wire Best Practices](https://github.com/google/wire/blob/master/docs/best-practices.md)  
[Google Wire FAQ](https://github.com/google/wire/blob/master/docs/faq.md)
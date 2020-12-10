# design_pattern
go 设计模式

[toc]
## 七大原则
### 开闭原则
开闭原则（Open-Closed Principle, OCP）规定“软件中的对象（类，模块，函数等等）应该对于扩展是开放的，但是对于修改是封闭的”，这意味着一个实体是允许在不改变它的原（源）代码的前提下变更它的行为。

### 依赖倒置原则（依赖反转原则）
依赖倒置原则（依赖反转原则 Dependency inversion principle，DIP）是指一种特定的解耦（传统的依赖关系创建在高层次上，而具体的策略设置则应用在低层次的模块上）形式，使得高层次的模块不依赖于低层次的模块的实现细节，依赖关系被颠倒（反转），从而使得低层次模块依赖于高层次模块的需求抽象。

原则规定：
- 高层次的模块不应该依赖于低层次的模块，两者都应该依赖于抽象接口。
- 抽象接口不应该依赖于具体实现。而具体实现则应该依赖于抽象接口。

### 单一职责原则
单一职责（单一功能原则 Single responsibility principle）规定每个类都应该有一个单一的功能，并且该功能应该由这个类完全封装起来。所有它的（这个类的）服务都应该严密的和该功能平行（功能平行，意味着没有依赖）。

单一职责原则是高内聚低耦合的指导方针，可以降低类的复杂度，提高类的可读性，提高系统的可维护性、降低变更引起的风险。其实，通俗来理解，一个类不能太累。我们在传统的软件工程，或者老旧的系统中经常能够看到，一个代码上万行的类。这种情况可能是由于历史原因导致的，但是不可否认，维护起来难度简直太大。因此，我们在软件设计过程中，要尝试将职责进行分离，不通职责封装在不同的类中，这样就能够降低我们设计软件的复杂度了。

### 接口隔离原则
接口隔离原则（interface-segregation principles，ISP）指明客户（client）不应被迫使用对其而言无用的方法或功能。接口隔离原则(ISP)拆分非常庞大臃肿的接口成为更小的和更具体的接口，这样客户将会只需要知道他们感兴趣的方法。这种缩小的接口也被称为角色接口（role interfaces）。接口隔离原则(ISP)的目的是系统解开耦合，从而容易重构，更改和重新部署。接口隔离原则是在SOLID (面向对象设计)中五个面向对象设计(OOD)的原则之一，类似于在GRASP (面向对象设计)中的高内聚性。

### 迪米特法则
迪米特法则（Law of Demeter，LOD）又叫作最少知识原则（The Least Knowledge Principle），一个类对于其他类知道的越少越好，就是说一个对象应当对其他对象有尽可能少的了解,只和朋友通信，不和陌生人说话。
一个软件实体应当尽可能少的与其他实体发生相互作用。这样，当一个模块修改时，就会尽量少的影响其他的模块，扩展会相对容易。迪米特法则是对软件实体之间通信的限制，它对软件实体之间通信的宽度和深度做出了要求。迪米特的其它表述方式为：
- 只与你直接的朋友们通信。
- 不要跟“陌生人”说话。
- 每一个软件单位对其他的单位都只有最少的知识，而且局限于那些与本单位密切相关的软件单位。

就像收银员和顾客的关系，收银员只管收到收到足够的钱然后给顾客商品，他不关心顾客还有多少钱。

### 里氏代换原则
里氏代换原则(Liskov Substitution Principle LSP)：任何基类可以出现的地方，子类一定可以出现。 LSP是继承复用的基石，只有当衍生类可以替换掉基类，软件单位的功能不受到影响时，基类才能真正被复用，而衍生类也能够在基类的基础上增加新的行为。里氏代换原则是对“开-闭”原则的补充。实现“开-闭”原则的关键步骤就是抽象化。而基类与子类的继承关系就是抽象化的具体实现，所以里氏代换原则是对实现抽象化的具体步骤的规范。

特别的，基类出现的地方可由衍生类替换，但衍生类出现的地方不可适用基类替换；子类可以实现父类的抽象方法，但是不能覆盖父类的非抽象方法；子类中可以增加自己特有的方法；当子类方法重载父类方法时，方法的前置条件(即方法的入参、输入)要比父类方法的输入参数更宽松或相等；当子类的方法实现父类的方法时(重写/重载或实现抽象方法)，方法的后置条件(即方法的输出/返回值)要比父类更严格或相等。

### 合成复用原则
合成复用原则（Composite Reuse Principle，CRP）又叫组合/聚合复用原则（Composition/Aggregate Reuse Principle，CARP）。它要求在软件复用时，要尽量先使用组合或者聚合等关联关系来实现，其次才考虑使用继承关系来实现。

如果要使用继承关系，则必须严格遵循里氏替换原则。合成复用原则同里氏替换原则相辅相成的，两者都是开闭原则的具体实现规范。


通常类的复用分为继承复用和合成复用两种，继承复用虽然有简单和易实现的优点，但它也存在以下缺点。
- 继承复用破坏了类的封装性。因为继承会将父类的实现细节暴露给子类，父类对子类是透明的，所以这种复用又称为“白箱”复用。
- 子类与父类的耦合度高。父类的实现的任何改变都会导致子类的实现发生变化，这不利于类的扩展与维护。
- 它限制了复用的灵活性。从父类继承而来的实现是静态的，在编译时已经定义，所以在运行时不可能发生变化。

采用组合或聚合复用时，可以将已有对象纳入新对象中，使之成为新对象的一部分，新对象可以调用已有对象的功能，它有以下优点。
- 它维持了类的封装性。因为成分对象的内部细节是新对象看不见的，所以这种复用又称为“黑箱”复用。
- 新旧类之间的耦合度低。这种复用所需的依赖较少，新对象存取成分对象的唯一方法是通过成分对象的接口。
- 复用的灵活性高。这种复用可以在运行时动态进行，新对象可以动态地引用与成分对象类型相同的对象。


## 创建型模式
- [简单工厂模式](https://github.com/hunyxv/design_pattern/tree/master/01%20-%20simplefactory)
- [工厂模式](https://github.com/hunyxv/design_pattern/tree/master/02%20-%20factorymethod)
- [抽象工厂模式](https://github.com/hunyxv/design_pattern/tree/master/03%20-%20abstractfactory)

## 结构型模式
- [适配器模式](https://github.com/hunyxv/design_pattern/tree/master/04%20-%20adapter)
- [装饰者模式&装饰器](https://github.com/hunyxv/design_pattern/tree/master/05%20-%20decorator)
- [外观模式](https://github.com/hunyxv/design_pattern/tree/master/06%20-%20facade)
- [组合模式](https://github.com/hunyxv/design_pattern/tree/master/07%20-%20composite)
- [享元模式](https://github.com/hunyxv/design_pattern/tree/master/08%20-%20flyweight)
- [桥接模式](https://github.com/hunyxv/design_pattern/tree/master/09%20-%20bridge)
- [代理模式](https://github.com/hunyxv/design_pattern/tree/master/10%20-%20proxy)

## 行为型模式
- [责任链模式](https://github.com/hunyxv/design_pattern/tree/master/11%20-%20chain_of_responsibility)
- [命令模式](https://github.com/hunyxv/design_pattern/tree/master/12%20-%20command)
- [观察者模式](https://github.com/hunyxv/design_pattern/tree/master/13%20-%20observer)
- [解释器模式](https://github.com/hunyxv/design_pattern/tree/master/14%20-%20interpreter)
- [状态模式](https://github.com/hunyxv/design_pattern/tree/master/15%20-%20state)
- [策略模式](https://github.com/hunyxv/design_pattern/tree/master/16%20-%20strategy)
- [模板模式](https://github.com/hunyxv/design_pattern/tree/master/17%20-%20template)
- [迭代器模式](https://github.com/hunyxv/design_pattern/tree/master/18%20-%20iterator)
- [中介者模式](https://github.com/hunyxv/design_pattern/tree/master/19%20-%20mediator)
- [备忘录模式](https://github.com/hunyxv/design_pattern/tree/master/20%20-%20memento)

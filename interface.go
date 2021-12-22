
// 接口
/*
	基本介绍
		interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量。到
		某个自定义类型(比如结构体Phone)要使用的时候，再根据具体情况把这些方法写出来(实现)。
			type Usb interface {
				Start()
				Stop()
			}
		
	基本语法
		type 接口名 interface {
			method1(参数列表) 返回值列表
			method2(参数列表) 返回值列表
		}

		实现接口所有方法：
		func(t 自定义类型) method1(参数列表) 返回值列表 {
			// 方法实现
		}
		func(t 自定义类型) method2(参数列表) 返回值列表 {
			// 方法实现
		}

	小结说明：
		1）接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计
		的多态和高内聚低耦合的思想。
		2）Go中的接口，不需要显示的实现。只要一个变量，含有接口类型中的所有方法，那么这个变
		量就实现了这个接口。因此，Go中没有implement这样的关键字。

	接口的注意事项和细节
		1）接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量(实例)。

		2）接口中所有的方法都没有方法体，即都是没有实现的方法。

		3）在Go中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该
		接口。

		4）一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋给接口类型。

		5）只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型。

		6）一个自定义类型可以实现多个接口。

		7）Go接口中不能有任何变量。

		8）一个接口(比如A接口)可以继承多个别的接口(比如B,C接口)，这时如果要实现A接口，也必须将
		B，C接口的方法也全部实现。
 
		9）interface类型默认是一个指针(引用类型)，如果没有对interface初始化就使用，那么会输出nil。

		10）空接口interface{}没有任何方法，所以所有类型都实现了空接口。

	接口的最佳实践
		实现sort.Sort(data interface)

	实现接口vs继承
		1）当A结构体继承了B结构体，那么A结构体就自动的继承了B结构体的字段和方法，并且可以直接使用。

		2）当A结构体需要扩展功能，同时不希望去破坏继承关系，则可以实现某个接口即可，因此我们可以
		认为：实现接口是对继承机制的补充。

		3）接口和继承解决的问题不同：
			继承的价值主要在于：解决代码的复用性和可维护性。
			接口的价值主要在于：设计，设计好各种规范(方法)，让其它自定义类型去实现这些方法。

		4）接口比继承更加灵活
			接口比继承更加灵活，继承是满足is - a的关系，而接口只需要满足like - a的关系。

		5）接口在一定程度上实现代码解耦。 
*/ 
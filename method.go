
// 方法
/*
	基本介绍
		Golang中的方法是作用在指定的数据类型上的（即：和指定的数据类型绑定），因此自定义类型，
		都可以有方法，而不仅仅是struct

	方法的声明和调用
		type Person struct{
			Name string
		}

		func (p Person)test() {
			fmt.Println(a.Name)
		}

		1）func (a A)test() {} 表示A结构体有一方法，方法名为test
		2）(a A)体现test方法是和A类型绑定的

	总结
		1）test方法和Person类型绑定
		2）test方法只能通过Person类型的变量来调用，而不能直接调用，也不能使用其它类型变量来调用
		3）func(p Person)test {}...p表示哪个Person变量调用，这个p就是它的副本，这点和函数传参非常相似
		4）p这个名字，有程序员指定，不是固定，比如修改成person也是可以的
		5）方法调用p.test()

	方法的调用和传参机制原理：
		说明：方法的调用和传参机制和函数基本一样，不一样的地方是方法调用时，会将调用方法的变量，当做实参
		也传递给方法
			1）在通过一个变量和调用方法时，其调用机制和函数一样
			2）不一样的地方是，变量调用方法时，该变量本身也会作为一个参数传递到方法（如果变量是值类型，
			则进行值拷贝，如果变量是引用类型，则进行地址拷贝）

	方法的声明（定义）
		func (receiver type) methodName (参数列表) (返回值列表) {
			方法体
			return 返回值
		}

		1）参数列表：表示方法输入
		2）receiver type：表示这个方法和type这个类型进行绑定，或者说该方法作用于type类型
		3）receiver type：type可以是结构体，也可以是其它的自定义类型
		4）receiver：就是type类型的一个变量（实例），比如：Person结构体的一个变量（实例）
		5）返回值列表：表示返回的值，可以多个
		6）方法主体：表示为了实现某一功能代码块
		7）return语句不是必须的
	
	方法的注意事项和细节
		1）结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式
		2）如程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理为了提高效率，通常我们将
		方法和结构体的指针类型绑定，编译器底层做了优化，(&c).area()等价于c.area()，因为编译器会自动加上&c
		3）Go中的方法作用在指定的数据类型上的（即：和指定的数据类型绑定），因此自定义类型，都可以有方法，而
		不仅仅是struct，比如int，float32等都可以有方法
		4）方法的访问范围控制的规则，和函数一样。方法名首字母小写，只能在本包访问，方法首字母大写，可以在本
		包和其他包访问
		5）如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出
	
*/
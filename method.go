
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
*/
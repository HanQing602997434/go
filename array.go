
//  数组
/*
	数组介绍
		数组可以存放多个同一类型数据。数组也是一种数据类型，在Go中，数组是值类型。

	定义数组
		var 数组名 [数组大小]数据类型
		var intArr [5]int

	使用：
		1.数组的地址可以通过数组名来获取&intArr

		2.数组的第一个元素的地址，就是数组的首地址

		3.数组的各个元素的地址间隔是依据数组的类型决定，比如int64->8 / int32->4

		4.四种初始化数组的方式
			var numArr01 [3]int = [3]int{1, 2, 3}

			var numArr02 = [3]int{5, 6, 7}

			var numArr03 = [...]int{8, 9, 10}

			var numArr04 = [...]int{1: 800, 0: 900, 2: 999} 
			// 类型推导
			strArr05 := [...]string{1: "timo", 0: "haha", 2: "luffy"}

	数组的遍历：
		1.常规遍历
			for i := 0; i < len(arr); i++ {

			}

		2.for-range结构遍历
			这是Go语言一种独有的结构，可以用来遍历访问数组的元素
			for index, value := range arr {

			}
			说明：
				第一个返回值index是数组的下标
				第二个value是该下标位置的值
				他们都是仅在for循环内部可见的局部变量
				遍历数组元素的时候，如果不像使用下标index，可以直接把下标index标为下划线_
				index和value的名称不是固定的，即程序员可以自行指定，一般命令为index和value

	数组使用注意事项和细节
		1.数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化

		2.var arr []int 这时arr就是一个slice切片

		3.数组中的元素可以是任意数据类型，包括值类型和引用类型，但是不能混用

		4.数组创建后，如果没有赋值，有默认值（零值）
			数值类型数组：默认值为0
			字符串类型数组：默认值为""
			bool类型数组：默认值为false

		5.使用数组的步骤1.声明数组并开辟空间 2给数组各个元素赋值（默认零值） 3 使用数组

		6.数组的下标是从0开始的

		7.数组下标必须在指定范围内使用，否则报panic：数组越界

		8.Go的数组属值类型，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响

		9.如想在其他函数中，去修改原来的数组，可以使用引用传递（指针方式）

		10.长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度
*/

// 切片
/*
	基本介绍
		1.切片得英文是slice

		2.切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制

		3.切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度len(slice)都一样

		4.切片的长度是可以变化的，因此切片是一个可以动态变化的数组

		5.切片定义的基本语法
			var切片名 []类型
			比如：var a []int

	快速入门
		var intArr [5]int = [...]int{1, 22, 33, 66, 99}
		
		//1. slice就是切片名
		//2. intArr[1:3]表示slice引用到intArr这个数组
		//3. 引用intArr数组的起始下标为1，最后的下标为3（但是不包含3）
		slice := intArr[1:3]
		
		fmt.Println("intArr =", intArr)
		fmt.Println("slice 的元素是 =", slice)
		fmt.Println("slice 的元素个数 =", len(slice))
		fmt.Println("slice 的容量 =", cap(slice)) // 切片的容量是可以动态变化的

	切片的原理
		1.slice是一个引用类型

		2.slice从底层来说，其实就是一个数据结构
		type slice struct {
			ptr *[2]int
			len int
			cap int
		}
*/
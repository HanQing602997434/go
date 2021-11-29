
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

	切片的使用
		第一种方式：定义一个切片，然后让切片去引用一个已经创建好的数组

		第二种方式：通过make来创建切片
			基本语法：var 切片名 []type = make([]type, len, cap)
			var slice []int = make([]int, 4, 10)
			1）通过make方式创建切片可以指定切片的大小和容量
			2）如果没有给切片的各个元素赋值，那么就会使用默认值[int = 0, float = 0, string = "", bool = false]
			3）通过make方式创建的切片对应的数组是由make底层维护，对外不可见，即只能通过slice去访问各个元素

		第三种方式：定义一个切片，直接就指定具体数组，使用原理类似make的方式
			var strSlice []string = []string{"tom", "jack", "mary"}

		方式一和方式二的区别
			方式一是直接引用数组，这个数组是事先存在的，程序员可见的，通过数组和切片都可以访问
			方式二是通过make来创建切片的，make也会创建一个数组，是由切片在底层进行维护，程序员是看不见的，只能通过切片访问

	切片的遍历
		for i := 0; i < len(slice); i++ {

		}

		for index, value := range(slice) {

		}

	切片的注意事项和细节说明
		1.切片初始化时 var slice = arr[startIndex:endIndex]
			说明：从arr数组下标为startIndex，到取到下标为endIndex的元素（不含arr[endIndex]）

		2.切片初始化时，仍然不能越界。范围在[0-len(arr)]之间，但是可以动态增长
			1）var slice = arr[0:end]可以简写 var slice = arr[:end]
			2）var slice = arr[start:len(arr)]可以简写 var slice = arr[start:]
			3）var slice = arr[0:len(arr)]可以简写 var slice = arr[:]

		3.cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素。

		4.切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用

		5.切片可以继续切片
*/
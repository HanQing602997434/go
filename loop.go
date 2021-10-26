
// 循环
/*
	for循环控制
		基本语法：
		for 循环变量初始化；循环条件；循环变量迭代 {
			循环操作（语句）
		}

		for i := 1; i <= 10; i++ {
			fmt.Println("hello world!")
		}

		细节：
			1）循环条件是返回一个布尔值的表达式
			2）for循环的第二种使用方式
				
				for 循环判断条件 {
					// 循环执行语句
				}
				将变量初始化和变量迭代写到其他位置

				j := 1
				for j <= 10 {
					fmt.Println("hello world~")
					j++
				}

			3）for循环的第三种使用方式
				
				for ; ; {
					// 循环执行语句
				}
				死循环

			4)for-range

				for index, val := range str {
				}
				使用的是字符的方式遍历
*/
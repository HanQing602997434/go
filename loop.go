
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

	while循环控制
		Go中没有while
		for {
			if i > 10 {
				break
			}
			fmt.Println("hello world!", i)
			i++
		}

	do while循环控制
		Go中没有do while
		for {
			fmt.Println("Hello World!", i)
			i++
			if i > 10 {
				break
			}
		}
		
	多重循环控制
		1）将一个循环放在另一个循环体内，就形成了嵌套循环。在外边的for称为外层循环	
		在里面的for循环称为内层循环。【建议一般使用两层，最多不要超过3层】
		2）实质上，嵌套循环就是把内层循环当成外层循环的循环体。当只有内层循环的循环
		条件为false时，才会完全跳出内层循环，才可结束外层的档次循环，开始下一次循环。
		3）外层循环次数为m次，内层为n次，则内层循环体实际上需要执行m*n=mn次
*/
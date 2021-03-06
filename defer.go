
// defer
/*
	在函数中，经常需要创建资源（比如：数据库连接、文件句柄、锁等），为了在函数执行完毕后，
	及时的释放资源，Go的设计者提供defer（延时机制）。

	细节：
		1.当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈中（defer栈），
		然后继续执行函数下一个语句。
	
		2.当函数执行完毕后，再从defer栈中，按照先入后出的方式出栈执行

		3.在defer将语句放入到栈时，也会将相关的值拷贝同时入栈。

	defer的最佳实践：
		defer最主要的价值是在，当函数执行完毕后，可以及时的释放函数创建的资源。看下模拟代码：
		func test() {
			// 关闭文件资源
			file = openfile(文件名)
			defer file.close()
			//其他代码...可以继续使用file句柄
		}

		func test() {
			// 释放数据库资源
			connect = openDatabase()
			defer connect.close()
			// 其他代码...可以继续使用connect
		}

		1.在Go编程中的通常做法是，创建资源后，比如（打开了文件，获取了数据库的连接，或者是锁资源）
		可以执行defer file.Close() defer connect.Close()。
		2.在defer后，可以继续使用创建资源。
		3.在函数完毕后，系统会依次从defer栈中，取出语句，关闭资源。
		4.这种机制，非常简洁，程序员不用再为在什么时机关闭资源而烦心。
*/
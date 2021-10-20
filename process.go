
// 流程控制
/*
	三大流程控制：
		1.顺序控制
			程序从上到下逐行执行，中间没有任何判断和跳转

		2.分支控制
			让程序有选择的执行，分三种
				1）单分支
					if 条件表达式 {
						执行代码块
					}

					Go支持在if语句中定义一个变量

				2）双分支
					if 条件表达式 {
						执行代码块
					} else {
						执行代码块
					}

					双分支只会执行其中的一个分支

				3）多分支控制
					if 条件表达式1 {
						执行代码块1
					} else if 条件表达式2 {
						执行代码块2
					} 
					... else {
						执行代码块n
					}

					多分支最多只会执行一个分支
					if 后面的表达式不允许是一个赋值语句

				4）嵌套分支
					if 条件表达式 {
						if 条件表达式 {
							if ...
						}
					}

		3.循环控制
*/
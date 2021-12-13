
// 方法和函数的却别
/*
	1.调用方式不一样
		函数的调用方式：	函数名（实参列表）
		方法的调用方式：	变量名.方法名（实参列表）

	2.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然

	3.对于方法（如struct的方法），接收者为值类型时，可以直接使用指针类型的变量
	调用方法，反过来同样可以。方法的调用主要看实参是什么类型，如果参数是值类型
	则进行值拷贝，如果参数是地址类型，则进行地址拷贝
*/

// init函数
/*
	介绍：
		每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被Go运行框架调用，
		也就是说init会在main函数前被调用。通常可以在init函数中完成初始化工作。
	
	init函数的注意事项和细节
		1）如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程全局变量定义->
		init函数->main函数

		2）init函数最主要的作用，就是完成一些初始化的工作

		3）如果main.go和utils.go都含有变量定义，init函数时，执行的流程utils.go的变量定义->
		init函数->main.go的变量定义->init函数->main函数
*/
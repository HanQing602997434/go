
// Go快速入门
/*
	windows下开发步骤：
		1.安装vscode

		2.go文件以.go

		3.go build编译.go文件，生成.exe文件

		4.go run运行源码，速度慢

	Go程序开发注意事项：
		1.go源文件以"go"为扩展名

		2.go应用程序的执行入口是main函数

		3.go语言严格区分大小写

		4.go方法由一条条语句构成，每个语句后不需要分号（go语言会在每行后自动加分号），
		这也体现了go的简洁性

		5.go编译器是一行行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同
		一行，否则报错

		6.go语言定义的变量或者import的包如果没有使用到，代码编译不能通过

		7.大括号都是成对出现，缺一不可

	Go转义字符(escape char)：
		\t：一个制表符，实现对齐的功能，通常用来排版
		\n：换行符
		\\：一个\
		\"：一个"
		\r：一个回车 fmt.Println("五虎上将\r子龙")

	Go开发常见问题和解决方法
		The system cannot find the file specified

		'hello.exe'不是内部或外部命令，也不是可运行的程序或批处理文件

		解决方案：源文件名写错或者不存在，或者当前路径错误

	Go代码快捷键：
		选中代码，按shift + tab键左移，按tab键右移

		命令gofmt main.go格式修正，gofmt -w main.go修正源文件

	Go代码风格，要求函数的左大括号和函数名同一行

		Go的设计思想：一个问题尽量只有一个解决方案

		Go一行代码尽量不超过80个字符，超过用逗号(,)换行

	DOS命令：
		md：创建文件夹
		dir：查看当前目录
		(cd /d f:)：切换到某个盘
		cd \：去根目录
		rd：删除空目录
		rd /q/s：不用询问删除目录
		echo hello > d:\abc.txt：向d盘abc.txt文件写入hello，没有abc.txt则创建
		copy d:\abc.txt f:\：将d盘abc.txt文件拷贝到f盘
		move d:\abc.txt f:\：将d盘abc.txt文件移动到f盘	
		del d:\abc.txt：删除d盘abc.txt文件
		cls：清屏
		exit：退出		
*/
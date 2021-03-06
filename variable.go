
// 变量
/*
	使用：

		// 变量声明
		var i int

		// 变量赋值
		i = 10

		// 变量使用
		fmt.Println("i=", i)
	
	注意事项：

		1.变量表示内存中的一个存储区域

		2.该区域有自己的名称（变量名）和类型（数据类型）

		3.Go变量使用的三种方式：
			1）指定变量类型，声明后若不赋值，使用默认值，int的默认值是0
			var i int
			
			2）根据值自行判定变量类型（类型推导）
			var num = 10.10

			3）省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误

		4.多变量声明

			1）同一种变量
				var n1, n2, n3 int

			2）不同种变量
				var n1, name, n3 = 100, "tom", 888

			3）多变量声明类型推导
				n1, name, n3 := 100, "tom", 888

			4）全局变量多变量声明
				var (
					n1 = 100
					n2 = 200
					name = "mary"
				)
				
			5）该区域内的数据值可以在同一类型范围内不断变化
				i = 10
				i = 20
				不允许i = 1.2

			6）变量在同一个作用域（在一个函数或代码块）内不能重名

			7）变量=变量名+值+数据类型，变量三要素

			8）变量如果没有赋初值，编译器会使用默认值，比如int默认值为0，string默认值为空串，小数默认值为0

		5.数据类型
			
				1.基本类型
					1）数值型：
						整数类型（int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, byte）
							fmt.Printf()可以用于做格式化输出。
							unsafe.Sizeof(n1)是unsafe包的一个函数，可以返回变量占用的字节数
							Go中int的使用遵守保小不保大的原则，在保证程序正确运行下，尽量使用空间小的数据类型。如年龄

						浮点类型（float32, float64）
							浮点型默认是float64
							浮点类型的大小不受操作系统影响
							科学计数法：5.1234e2 = 5.1234 * 10的2次方

					2）字符型（没有专门的字符型，使用byte来保存单个字母字符）
						字符常量是用单引号''括起来的单个字符
						Go中允许使用转义字符'\'来将其后的字符转变为特殊字符常量
						Go语言的字符使用UTF-8编码，英文字母占1个字节 汉字占3个字节
						在Go中，字符的本质是一个整数，直接输出时，是该字符对应的UTF-8编码的码值
						可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的unicode字符
						字符类型时可以进行运算的，相当于一个整数，因为它都对应有unicode码

					3）布尔型（bool）
						取值只有true和false

					4）字符串（string）
						字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。也就是说对于
						传统的字符串是由字符组成的，而Go的字符串不同，它是由字节组成的。
						如果我们保存的字符码值在byte范围内，可以用byte保存，如果超过byte范围，考虑用int保存
						字符串一旦赋值了，字符串就不能修改了，在Go中字符串是不可变的
						使用反引号``可以输出源码
						字符串拼接方式： +，+=
						当一个拼接的操作很长时，可以分行写，把+保留在上一行的结尾

				2.派生/复杂数据类型
					1）指针（Pointer）

					2）数组

					3）结构体（struct）

					4）管道（Channel）

					5）函数

					6）切片（slice）

					7）接口（interface）

					8）map

		6.基本数据类型的转换
			Golang和java/c不同，Go在不同类型的变量之间赋值需要显示转换。也就是说Golang中数据类型不能自动转换

			基本语法：
				表达式T(v)将值v转换成类型T

			细节：
				被转换的是变量存储的数据（即值），变量本身的数据类型并没有改变
				一个很大的值转换成很小的类型，结果是按溢出处理

		7.基本数据类型转换为string
			fmt.Sprintf()根据format参数生成格式化的字符串并返回该字符串
				int转string：str := fmt.Sprintf("%d", num1)
				float64转string：str = fmt.Sprintf("%f", num2)
				bool转string：str = fmt.Sprintf("%t", b)
				byte转string：str = fmt.Sprintf("%c", myChar)

			strconv.FormatInt()利用strconv包里的转换函数
				int转string：str = strconv.FormatInt(int64(num1), 10)
				float64转string：str = strconv.FormatFloat(num2, 'f', 10, 64)
				bool转string：str = strconv.FormatBool(b)

		8.string类型转换基本数据类型
			使用strconv包的函数
				string转bool：b, _ = strconv.ParseBool(str)
				string转int64：n1, _ = strconv.ParseInt(str2, 10, 64)
				string转float64：f1, _ = strconv.ParseFloat(str3, 64)
			
			string转基本数据类型如果不成功就将其转换为默认值
*/
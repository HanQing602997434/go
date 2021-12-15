
// 面向对象编程
/*
	步骤
		1）声明（定义）结构体，确定结构体名
		2）编写结构体字段
		3）编写结构体的方法

	创建结构体变量时指定字段值
		Golang在创建结构体实例（变量）时，可以直接指定字段的值

		方式1：
			var stu1 Student = Student{"tom", 10}

			stu2 := Student{"tom", 10}

			// 在创建结构体变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
			var stu3 Student = Student{
				Name : "mary",
				Age : 30,
			}

			stu4 := Student{
				Age : 20,
				Name : "mary",
			}

		方式2：返回结构体的指针类型(!!!)
			var stu5 = &Stu("小王", 29)
			stu6 := &Stu("小王~", 39)

			// 在创建结构体指针变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
			var stu7 = &Stu{
				Name : "小李"，
				Age : 49,
			}
			
			stu8 := &Stu{
				Age : 59,
				Name : "小李~",
			}
*/
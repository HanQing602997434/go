
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

			var stu3 Student = Student{
				Name : "mary",
				Age : 30,
			}

			stu4 := Student{
				Name : "mary",
				Age : 20,
			}

		方式2：
*/

// 类型断言
/*
	类型断言
		由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言

		需求：如何将一个接口变量，赋给自定义类型的变量 => 引出类型断言
		type Point struct {
			x int
			y int
		}
		func main() {
			var a interface{}
			var point Point = Point{1, 2}
			a = point // ok
			// 如何将a赋给一个Point变量？
			var b Point
			b = a // 可以吗？=>error
			fmt.Println(b)
		}

		func main() {
			var a interface{}
			var point Point = Point{1, 2}
			a = point
			var b Point
			b = a.(Point) // 类型断言
			fmt.Println(b)
		}
		b = a.(Point)就是类型断言，表示判断a是否指向Point类型的变量，如果是就转成Point
		类型并赋给b变量，否则报错。

		在进行类型断言时，如果类型不匹配，就会报panic，因此进行类型断言时，要确保原来的空
		接口指向的就是断言的类型。
		在进行类型断言时，带上检测机制，如果成功就ok，否则也不要报panic
*/
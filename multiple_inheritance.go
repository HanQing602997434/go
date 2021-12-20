
// 多重继承
/*
	结构体的匿名字段是基本数据类型
		type A struct {
			Name string
			Age int
		}

		type Stu struct {
			A
			int
		}
		通过stu.int = 90来使用

	多重继承
		type Goods struct {
			Name string
			Price float64
		}

		type Brand struct {
			Name string
			Address string
		}

		type TV struct {
			Goods
			Brand
		}

	多重继承的细节
		1）如嵌入的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体类型名来区分。
		2）为了保证代码的简洁性，建议大家尽量不要使用多重继承。
*/
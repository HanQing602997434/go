
// 结构体使用细节
/*
	1.结构体内存是连续的

	2.结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段（名字、个数和类型）

	3.结构体进行type重新定义（相当于取别名），Golang认为是新的数据类型，但是相互间可以强转

	4.struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
		type Monster struct{
			Name string `json:"name"`
			Age int `json:"age"`
			Skill string `json:"skill"`
		}
*/
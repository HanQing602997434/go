
// map使用细节和陷阱
/*
	细节
		1.map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后，会直接修改原来的map

		2.map的容量达到后，再想map增加元素，会自动扩容，并不会发生panic，也就是说map能动态的增长
		键值对（key-value）

		3.map的value也经常使用struct类型，更适合管理复杂的数据（比前面value是一个map更好），比如
		value为Student结构体
*/
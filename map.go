
// map
/*
	map介绍
		map是key-value数据结构，又称为字段或者关联数组。类似其他编程语言的集合

	基本语法
		var map变量名 map[keytype]valuetype

		key可以是什么类型
			golang中的map的key可以是很多种类型，比如bool、数字、string、指针、channel，
			还可以是只包含前面几个类型的接口、结构体、数组
			通常为int、string

			注意：slice、map还有function不可以作为key，因为这几个没法用==来判断

		value可以是什么类型
			valuetype的类型和key基本一样
			通常为数字、string、map、struct

		map声明的举例
			var a map[string]string
			var a map[string]int
			var a map[int]string
			var a map[string]map[string]string

			注意：声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用
			a = make(map[string]string, 10)

	map的使用方式
		方式1：
			// 先声明
			var cities map[string]string
			// 再make
			cities = make(map[string]string, 10)

		方式2：
			// 声明，就直接make
			var cities = make(map[string]string)

		方式3：
			var cities map[string]string = map[string]string{
				"no4" : "成都"，
			}
			cities["no1"] = "北京"

	map的增删改查操作
		map的增加和更新
			map["key"] = value // 如果key还没有，就是增加，如果存在就是修改

		map删除
			delete(map, "key")，delete是一个内置函数，如果key存在，就删除该key-value，
			如果key不存在，不操作，但是也不会报错

			细节说明
				1）如果我们要删除map的所有key，没有一个专门的方法一次删除，可以遍历一下，逐个删除

				2）或者map = make(...)，make一个新的，让原来的称为垃圾，被gc回收

		map查找
			val, ok = map["key"]
			if ok {
				fmt.Println("有key这个元素")
			} else {
				fmt.Println("没有key这个元素")
			}

	map的遍历
		map的遍历使用for-range的结构遍历
		for k, v := range map {
			
		}
*/
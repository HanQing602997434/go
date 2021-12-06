
// map排序
/*
	map的元素是无序的

	如果按照map的key的顺序进行排序输出
		1.先将map的key放入到切片中
			var keys []int
			for k, _ := range map1 {
				keys = append(keys, k)
			}
		
		2.对切片排序
			sort.Ints(keys)

		3.遍历切片，然后按照key来输出map的值
			for _, k := range keys {
				fmt.Println("map[%v] = %v\n", k, map1[k])
			}
*/
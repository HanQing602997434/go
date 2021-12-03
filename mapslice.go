
// map切片
/*
	基本介绍
		切片的数据类型如果是map，则我们称为slice of map，map切片，这样使用则map个数就可以动态变化了

	使用切片的append函数，可以动态的增加monster
		newMonster := map[string]string{
			"name" : "新的妖怪~火云邪神",
			"age" : "200",
		}

		monsters = append(monsters, newMonster)
*/
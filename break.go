
// 跳转控制break
/*
	随机数包：math/rand
	随机函数：rand.Intn(100)，生成[0, 100)的整数
	随机种子：rand.Seed(time.Now().UnixNano())秒、rand.Seed(time.Now().UnixNano())纳秒

	break是用来终止语句的执行，用于结束for循环，switch分支
	break出现再多层嵌套的语句块时，默认跳出最近的for循环，可以通过标签指定跳出那一层语句块
	标签使用：label1:
*/
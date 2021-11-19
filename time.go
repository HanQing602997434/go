
// 时间和日期相关函数
/*
	基本介绍
		1.

		2.time.Time类型，用于表示时间

		3.获取当前时间
			now := time.Now()

		4.通过now都可以获取年月日时分秒
			now.Year()
			int(now.Month()
			now.Day()
			now.Hour()
			now.Minute()
			now.Second()

		5.格式化日期和时间
			fmt.Printf("当前年月日 %02d-%02d-%02d %d:%02d:%02d\n", now.Year(),
			now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

			dateStr = fmt.Sprintf("当前年月日 %02d-%02d-%02d %d:%02d:%02d\n", now.Year(),
			now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

			fmt.Printf(now.Format("2021/11/18 18:18:18"))

		6.时间的常量
			const (
				Nanosecond Duration = 1 // 纳秒
				Microsecond = 1000 * Nanosecond // 微秒
				Millisecond = 1000 * Microsecond // 毫秒
				Second = 1000 * Millisecond // 秒
				Minute = 60 * Second // 分钟
				Hour = 60 * Minute // 小时
			)

		7.休眠
		func Sleep(d Duration)
*/
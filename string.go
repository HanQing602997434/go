
// 字符串
/*
	字符串在我们程序开发中，使用的非常多，常用的系统函数如下：
		1.统计字符串的长度，按字节 len(str)，Go的编码统一为utf-8，
		字母和数字占一个字节，汉字占3个字节
		
		2.字符串遍历，同时处理有中文的问题，先将字符串转成切片 r := []rune(str)

		3.字符串转整数 n, err := strconv.Atoi("123")，如果转换错误会返回一个err

		4.整数转字符串 str = strconv.Itoa(12345)

		5.字符串转[]byte var bytes = []byte("hello go")

		6.[]byte转字符串 str = string([]byte{97, 98, 99})

		7.10进制转成2、8、16进制 str = strconv.FormatInt(123, 2)

		8.查找子串是否在指定的字符串中 strings.Contains("seafood", "foo")

		9.统计一个字符串有几个指定的子串 strings.Count("ceheese", "e")

		10.不区分大小写的字符串比较（==是区分字母大小写）strings.EqualFold("abc", "Abc")

		11.返回子串在字符串第一次出现的index值，如果没有返回-1 strings.Index("NLT_abc", "abc")
*/
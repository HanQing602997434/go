
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

		12.返回子串在字符串最后一次出现的index，如没有返回-1 strings.LastIndex("go golang", "go")

		13.将指定的子串替换成另外一个子串 strings.Replace("go go hello", "go", "go语言", n) n可以
		指定你希望替换几个，如果 n = -1 表示全部替换

		14.按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组 strings.Split("hello,world,ok", ",")

		15.将字符串的字母进行大小写的转换 strings.ToLower("Go") strings.ToUpper("Go")

		16.将字符串左右两边的空格去掉 strings.TrimSpace("tn a lone gopher ntm   ")

		17.将字符串左右两边指定的字符去掉 strings.Trim("!hello!", "!")

		18.将字符串左边指定的字符去掉 strings.TrimLeft("!hello!", "!")

		19.将字符串右边指定的字符去掉 strings.TrimRight("!hello!", "!")

		20.判断字符串是否以指定的字符串开头 strings.HasPrefix("ftp://192.168.0.1", "ftp")

		21.判断字符串是否以指定的字符串结尾 strings.HasSuffix("NTL_abc.jpg", "abc")
*/
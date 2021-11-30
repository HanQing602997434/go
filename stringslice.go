
// string和slice
/*
	1.string底层是一个byte数组，因此string也可以进行切片处理

	2.string和切片在内存的形式类似

	3.string是不可变的，也就是说不能通过str[0] = 'z'方式来修改字符串

	4.如果需要修改字符串，可以先将string->[]byte / 或者[]rune ->修改->重写转成string
		细节，我们转成[]byte后，可以处理英文和数字，但是不能处理中文
		原因是 []byte 字节来处理，而一个汉字，是3个字节，因此就会出现乱码
		解决方法是将 string转成[]rune即可，因为[]rune是按字符处理，兼容汉字
*/
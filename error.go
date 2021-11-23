
// 错误处理
/*
	1.在默认情况下，当发生错误后（panic），程序就是退出（崩溃）

	2.如果我们希望：当发生错误后，可以捕获到错误，并进行处理，保证程序可以继续执行。
	还可以在捕获到错误后，给管理员一个提示（邮件、短信）

	3.Go语言追求简洁优雅，所以，Go语言不支持传统的try...catch...finally

	4.Go的引入的处理方式为：defer、panic、recover

	5.这几个异常的使用场景可以这么简单描述：Go种可以抛出一个panic的异常，然后在defer中
	通过recover捕获这个异常，然后正常处理
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err =", err)
			}
		}()
*/
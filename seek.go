
// 查找
/*
	介绍
		在Golang中，我们常用的查找有两种：
			顺序查找
				顺序遍历，效率低

			二分查找（数据有序），思路如下：
				1.数组是有序的，并且从小到大排序，我们要查找的数是findVal

				2.先找到中间的下标mid = (leftIndex + rightIndex) / 2，然后用中间下标的值和findVal进行比较
					1）如果arr[mid] > findVal，就应该向leftIndex -- (mid - 1)
					2）如果arr[mid] < findVal，就应该向(mid + 1) -- rightIndex
					3）如果arr[mid] == findVal，则找到了
					4）上面的1 2 3的逻辑会递归执行

				3.递归退出的条件
					if leftIndex > rightIndex {
						// 找不到
						return
					}
*/
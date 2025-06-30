package main

import "fmt"

// func function_name( [parameter list] ) [return_types] {
//    函数体
// }
/* 函数返回两个数的最大值 */
func max(num1 int, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 交换
func swap(x, y string) (string, string) {
	return y, x
}

func learningArray() {
	// var arrayName [size]dataType
	// var int5 [5]int
	// var str5 [5]string
	// var int5 = [5]int{1, 2, 3, 4, 5}
	// numbers := [5]int{1, 2, 3, 4, 5}

	// 数组长度不确定
	// var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	// balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//  将索引为 1 和 3 的元素初始化
	// balance := [5]float32{1:2.0,3:7.0}

	// 对下标 4 赋值
	// balance[40] = 50.0

	// 取值
	// var salary float32 = balance[9]

	// 多维数组
	// var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

	// a := [3][4]int{
	// 	{0, 1, 2, 3} ,   /*  第一行索引为 0 */
	// 	{4, 5, 6, 7} ,   /*  第二行索引为 1 */
	// 	{8, 9, 10, 11},   /* 第三行索引为 2 */
	// 	}

	// sites := [2][2]string{}
	// // 向二维数组添加元素
	// sites[0][0] = "Google"
	// sites[0][1] = "Runoob"
	// sites[1][0] = "Taobao"
	// sites[1][1] = "Weibo"

	// // 创建空的二维数组
	// animals := [][]string{}
	// // 创建三一维数组，各数组长度不同
	// row1 := []string{"fish", "shark", "eel"}
	// row2 := []string{"bird"}
	// row3 := []string{"lizard", "salamander"}
	// // 使用 append() 函数将一维数组添加到二维数组中
	// animals = append(animals, row1)
	// animals = append(animals, row2)
	// animals = append(animals, row3)

}

// Go 没有三目运算符, 不支持 ? : 形式的条件判断

// Go 语言最少有个 main() 函数
func main() {
	fmt.Println("func")

	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)

	fmt.Printf("最大值是 : %d\n", ret)

	str_a, str_b := swap("Google", "Runoob")
	fmt.Println(str_a, str_b)

	learningArray()

}

// Go 语言程序中全局变量与局部变量名称可以相同, 但是函数内的局部变量会被优先考虑

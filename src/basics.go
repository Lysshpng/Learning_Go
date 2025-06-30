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

}

// Go 语言程序中全局变量与局部变量名称可以相同, 但是函数内的局部变量会被优先考虑

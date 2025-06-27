// 表示一个可独立执行的程序, 每个Go应用程序都包含一个名为main的包
package main

// fmt 包实现了格式化I/O的函数
import "fmt"

// main函数是每一个可执行程序所必须包含的
// 在main之前可能还有一个init
func main() {
	// { 不能单独放在一行, 在运行时会产生错误
	fmt.Println("Hello, World!")
	// Println 带回车
	// Print 不带回车

	// 格式化字符串
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	// 使用Sprintf
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	// 使用Printf
	fmt.Printf(url, stockcode, enddate)
	fmt.Println()

	// 变量
	// var name typer
	var intA int
	intA = 1
	fmt.Println(intA)
	// var name typer = value
	var strB string = "test"
	fmt.Println(strB)
	// var name = value 根据值自行判定变量类型
	var boolC = true
	fmt.Println(boolC)
	// name := value
	floatC := 1.234
	fmt.Println(floatC)

	// 多变量声明
	// var name1, name2, name3 typer
	var name1, name2, name3 bool
	name1, name3 = true, false
	name4, name5 := true, false
	fmt.Println(name1, name2, name3, name4, name5)

	// 变量类型默认值
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

}

/**
 * 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1,
 * 那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包）,
 * 这被称为导出（像面向对象语言中的 public）;
 * 标识符如果以小写字母开头, 则对包外是不可见的,
 * 但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）
 */

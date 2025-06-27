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
	// %d 表示整型数字, %s 表示字符串
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
	// name := value 这是使用变量的首选形式, 但是它只能被用在函数体内, 而不可以用于全局变量的声明与赋值
	floatC := 1.234
	fmt.Println(floatC)

	// 多变量声明
	// var name1, name2, name3 typer
	var name1, name2, name3 bool
	name1, name3 = true, false
	name4, name5 := true, false
	fmt.Println(name1, name2, name3, name4, name5)

	var int1, int2 int
	var str3 string
	int1, int2, str3 = 5, 7, "abc"

	int4, bool5, str6 := 5, false, "str6"

	_, int7 := 6, 7

	fmt.Println(int1, int2, str3, int4, bool5, str6, int7)

	// 变量类型默认值
	var in int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", in, f, b, s)

	// 声明了一个局部变量却没有在相同的代码块中使用它, 会得到编译错误
	// 全局变量是允许声明但不使用的

	// 常量
	const const1, const2 = 11, "str"
	fmt.Println(const1, const2)

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	// iota, 特殊常量, 可以认为是一个可以被编译器修改的常量
	// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前), const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)
	// 第一个 iota 等于 0, 每当 iota 在新的一行被使用时, 它的值都会自动加 1
	const (
		it1 = iota
		it2 = iota
	)

	const (
		iota_a = iota //0
		iota_b        //1
		iota_c        //2
		iota_d = "ha" //独立值，iota += 1
		iota_e        //"ha"   iota += 1
		iota_f = 100  //iota +=1
		iota_g        //100  iota +=1
		iota_h = iota //7,恢复计数
		iota_i        //8
	)
	fmt.Println(iota_a, iota_b, iota_c, iota_d, iota_e, iota_f, iota_g, iota_h, iota_i)

	// k和l相当于是默认与前一个相同的操作
	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i=", i) // 1
	fmt.Println("j=", j) // 6
	fmt.Println("k=", k) // 12
	fmt.Println("l=", l) // 24
}

/**
 * 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头, 如：Group1,
 * 那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包）,
 * 这被称为导出（像面向对象语言中的 public）;
 * 标识符如果以小写字母开头, 则对包外是不可见的,
 * 但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）
 */

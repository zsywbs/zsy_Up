// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	fmt.Println("My favorite number is ", rand.Intn(10))
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	fmt.Println(math.Pi)
// }

// package main

// import "fmt"

// func add(x, y int) int {
// 	return x + y
// }
// func main() {
// 	fmt.Println(add(42, 13))
// }

// package main

// import "fmt"

// func swap(x, y string) (string, string) {
// 	return y, x
// }
// func main() {
// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// }

/* package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(split(17))
}
*/

/* package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
*/

/* package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
*/

// 简洁赋值语句:=可代替var
// package main

// import "fmt"

// func main() {
// 	var i, j int = 1, 2
// 	k := 3
// 	c, python, java := true, false, "no！"
// 	fmt.Println(i, j, k, c, python, java)
// }

/* package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type:	%T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type:  %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type:	%T Value: %v\n", z, z)

}
*/

/* package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string = "zsy_Up"
	fmt.Println("%v %v %v %q\n", i, f, b, s)
}
*/

/* package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	var f float64 = math.Sqrt(x*x + y*y)

	var z uint = f

	fmt.Println(x, y, z)
}
*/

/* package main

import "fmt"

func main() {
	v := 42.456 + 5i
	fmt.Printf("v is of type %T\n", v)
}
*/

/* package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Gp rules?", Truth)
} */

/* package main

import "fmt"

const (
	// 将1左移 100位来创建一个非常大的数字
	// 即这个数的二进制是1后面跟着的100个0
	Big = 1 << 100
	// 再往右移99位
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Big)) //超出整数范围了
}
*/
/*
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
*/

/* package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// for *
// 是go中的while
/* package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
*/

// 无限循环

/* package main

func main() {
	for {

	}
}
*/

/* package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
*/

// /* package main

// import (
// 	"fmt"
// 	"math"
// )

// // math.pow(x,y) 表示x的y次方
// func pow(x, n, lim float64) float64 {
// 	if v := math.Pow(x, n); v < lim {
// 		return v
// 	}
// 	return lim   // if和c中的for一样，所以v只能在if的作用域中
// }

// func main() {
// 	fmt.Println(
// 		pow(3, 2, 10),
// 		pow(3, 3, 20),
// 	)
// }
//  */

/* package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)	// v同样可以用于else语句中
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
} */

/* package main

import "fmt"

func getabs(x float64) float64 {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}

	return x
}

func Sqrt(x float64) float64 {
	z := 1.0
	if x < 0 {
		return 0
	} else if x == 0 {
		return 0
	} else {
		for getabs(z*z-x) > 1e-6 {
			z = (z + x/z) / 2
		}
		return z
	}
}

func main() {
	fmt.Println(Sqrt(2))
} */

/* package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")
	switch zsy := runtime.GOOS; zsy { // 存有疑问
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", zsy)
	}
} */

/* package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Sunday?")
	today := time.Now().Weekday()//获取当前是星期几
	switch time.Saturday + 1 {
	case today + 0:
		fmt.Println("Today. ")
	case today + 1:
		fmt.Println("Tomorrow. ")
	case today + 2:
		fmt.Println("In two days. ")
	default:
		fmt.Println("Too far away. ")
	}
}
*/

/* package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
*/

/* package main

import "fmt"

func main() {
	defer fmt.Println("World") // defer语句会将函数推迟到外层函数返回之后执行

	fmt.Print("hello")
} */

/* package main

import "fmt"

func main() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) //推迟的函数调用会被压入一个栈（先进后出）

	}
	fmt.Println("done")
} */

/* package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
*/
/*
package main

import "fmt"

type Vertex struct { // 变量，类型，种类
	X int
	Y int
}

func main() {
	// fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

} */

/* package main

import "fmt"

type Vertex struct {
	X, Y int
	// Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // x：1 将1赋值给x，Y被隐式赋值为0（name ：语法列出部分字段）
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	// v := Vertex{1, 2}
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)
	fmt.Println(v1, p, v2, v3)
}
*/
/*
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} //输出数组名就是输出数组全部数据
	fmt.Println(primes)
}
*/

/* package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //创建了一个切片，下表从1-3的元素
	fmt.Println(s)
}
*/

/* package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"  //更改切片元素会修改原来的数据
	fmt.Println(a, b)
	fmt.Println(names)
}


*/

/* package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)  //都是用[]括起来的
}
*/

/* package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s) //输出[3,5,7]  //输出数组s之后，s数组中只有输出的数组元素

	s = s[:2]
	fmt.Println(s) //输出[2,3]

	s = s[1:] //输出[3,5,7,11,13]
	fmt.Println(s)
}
*/

/* package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	//截取切片的长度使其为0
	s = s[:0]
	printSlice(s)

	//拓展其长度
	s = s[:4] //若长度超过容量会报错
	printSlice(s)

	//舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
} // 格式化字符串%v，%T
*/

// nil切片 ,切片的零值是nil

/* package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
*/

/* package main

import "fmt"

func main() {
	a := make([]int, 5) // 类型-长度 len（）
	printSlice("a", a)

	b := make([]int, 0, 5) // 类型-长度-容量 len（） cap（）
	printSlice("b", b)     // 有长度才能有值

	c := b[0:]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
} */

/* package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上X和0
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][1] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}			// strings.Join(board[i]) 什么方法？？？、
}
*/

/* package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0) //append的第一个参数是一种元素类型的切片，后面数据类型相同，则添加到该数组
	printSlice(s)

	// 这个切片按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3) // 添加五个元素的时候 会分配一个更大的数组，容量为6
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}
*/

/* package main

import "fmt"

var Pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range Pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}		// 会返回两个值，i：当前元素下标，v：当前元素
}
*/

/* package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // uint()方法？？？
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
} */

/* package main

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for x := range a {
		b := make([]uint8, dx)
		for y := range b {
			b[y] = uint8(x * y)
		}
		a[x] = b
	}
	return a
}

func main() {
	pic.Show(Pic)
}
*/

/* package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
} */

/* package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

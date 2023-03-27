package main

func main() {

	//num1 := 2
	//num2 := 3
	//res1 := calculate.Sum(num1, num2)
	//fmt.Println("res1", res1)
	//res2 := calculate.Mul(num1, num2)
	//fmt.Println("res2", res2)

	//通过:=同时定义多个变量, 只要任意一个变量没有定义过,都会做退化赋值操作
	//// 変数 num1 を定義します
	//num1 := 10
	//// num2が定義されていないため、num1とnum2の2つの変数を同時に定義します。
	//// そのため、num1の:=は代入演算子に簡略化されますが、num2の:=は定義と代入を同時に行います。
	//num1, num2 := 20, 30
	//fmt.Println("num1 = ", num1)
	//fmt.Println("num2 = ", num2)

	//num1 := 10
	//num2 := 20
	//// 报错, 因为num1,和num2都已经被定义过
	//// 至少要有任意一个变量没有被定义过,才会退化赋值
	//num1, num2 := 30, 40
	//fmt.Println("num1 = ", num1)
	//fmt.Println("num2 = ", num2)

	//Go语言数值类型之间转换
	//var num0 int = 10
	//var num1 int8 = 20
	//var num2 int16
	////num2 = num0 // Cannot use 'num0' (type int) as the type int16
	////num2 = num1 // Cannot use 'num1' (type int8) as the type int16
	//num2 = int16(num0)
	//num2 = int16(num1)
	//fmt.Println(num2)
	//
	//var num3 float32 = 3.14
	//var num4 float64
	////num4 = num3 // Cannot use 'num3' (type float32) as the type float64
	//num4 = float64(num3)
	//fmt.Println(num4)
	//
	//var num5 byte = 11
	//var num6 uint8
	//num6 = num5 // 这里不是隐式转换, 不报错的原因是byte的本质就是uint8
	//fmt.Println(num6)
	//
	//var num7 rune = 11
	//var num8 int32
	//num8 = num7 // 这里不是隐式转换, 不报错的原因是rune的本质就是int32
	//fmt.Println(num8)

	//数值类型和字符串类型之间转换
	//var num1 int32 = 65
	//var str1 string = string(num1)
	//fmt.Println(str1) // A  可以将整型强制转换, 但是会按照ASCII码表来转换但是不推荐这样使用
	//
	//var num2 float32 = 3.14
	//var str2 string = string(num2) // Cannot convert an expression of the type 'float32' to the type 'string'
	//fmt.Println(str2)
	//
	//var str3 string = "97"
	//var num3 int = int(str3) // Cannot convert an expression of the type 'string' to the type 'int'
	//fmt.Println(num3)

	//数值类型转字符串类型  strconv..FormatXxx()
	//var num1 int32 = 10
	//str1 := strconv.FormatInt(int64(num1), 10)
	//fmt.Println(str1) // 10
	//
	//str2 := strconv.FormatInt(int64(num1), 2)
	//fmt.Println(str2) // 1010
	//
	//var num5 float64 = 3.1234567890123456789
	//// 第一个参数: 需要转换的实型, 必须是float64类型
	//// 第二个参数: 转换为什么格式,f小数格式, e指数格式
	//// 第三个参数: 转换之后保留多少位小数, 传入-1按照指定类型有效位保留
	//// 第四个参数: 被转换数据的实际位数,float32就传32, float64就传64
	//// 将float64位实型,按照小数格式并保留默认有效位转换为字符串
	//str3 := strconv.FormatFloat(num5, 'f', -1, 32)
	//fmt.Println(str3) // 3.1234567
	//str4 := strconv.FormatFloat(num5, 'f', -1, 64)
	//fmt.Println(str4) // 3.1234567890123457
	//str5 := strconv.FormatFloat(num5, 'f', 2, 64)
	//fmt.Println(str5) // 3.12
	//str6 := strconv.FormatFloat(num5, 'e', 2, 64)
	//fmt.Println(str6) // 3.12e+00
	//
	//var num6 bool = true
	//str7 := strconv.FormatBool(num6)
	//fmt.Println(str7)

	// 字符串类型转数值类型  strconv.ParseXxx()
	//var str1 string = "125"
	//// 第一个参数: 需要转换的数据
	//// 第二个参数: 转换为几进制
	//// 第三个参数: 转换为多少位整型
	//// 注意点: ParseInt函数会返回两个值, 一个是转换后的结果, 一个是错误
	//// 如果被转换的数据转换之后没有超出指定的范围或者不能被转换时,
	//// 那么错误为nil, 否则错误不为nil
	//// 将字符串"125"转换为10进制的int8
	//num1, err := strconv.ParseInt(str1, 10, 8)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(num1) // 125
	//
	//var str2 string = "150"
	//num2, err := strconv.ParseInt(str2, 10, 8)
	//if err != nil {
	//	fmt.Println(err) // strconv.ParseInt: parsing "150": value out of range
	//}
	//fmt.Println(num2) // 127
	//
	//var str3 string = "3.1234567890123456789"
	//num3, err := strconv.ParseFloat(str3, 32)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(num3) // 3.1234567165374756
	//
	//var str4 string = "true"
	//num4, _ := strconv.ParseBool(str4)
	//fmt.Println(num4) // true

	// 字符串类型转换为数值类型时,如果不能转换除了返回error以外,还会返回对应类型的默认值
	//var str1 string = "abc"
	//num1, _ := strconv.ParseInt(str1, 10, 32)
	//fmt.Println(num1) // 0
	//
	//num2, _ := strconv.ParseFloat(str1, 32)
	//fmt.Println(num2) // 0
	//
	//num3, _ := strconv.ParseBool(str1)
	//fmt.Println(num3) // false

	// 字符串类型和整型快速转换
	//var num1 int32 = 110
	//// 快速将整型转换为字符串类型
	//// 注意:Itoa方法只能接受int类型
	//var str1 string = strconv.Itoa(int(num1))
	//fmt.Println(str1) // 110
	//
	//var str2 string = "666"
	//// 快速将字符串类型转换为整型
	//// 注意: Atoi方法返回两个值, 一个值是int,一个值是error
	//// 如果字符串能被转换为int,那么error为nil, 否则不为nil
	//num2, err := strconv.Atoi(str2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(num2) // 666

	// 数值类型转字符串类型其它方式
	//var num1 int32 = 110
	//// Sprintf函数和Printf函数很像, 只不过不是输出而将格式化的字符串返回给我们
	//var str1 string = fmt.Sprintf("%d", num1)
	//fmt.Println(str1) // 110
	//
	//var num2 float32 = 3.14
	//var str2 string = fmt.Sprintf("%f", num2)
	//fmt.Println(str2) // 3.140000
	//
	//var num3 bool = true
	//var str3 string = fmt.Sprintf("%t", num3)
	//fmt.Println(str3) // true

	// 猜数字游戏
	//// 设置随机种子
	//rand.Seed(time.Now().Unix())
	//
	//// 随机生成一个1-100之间的数字
	//number := rand.Intn(100) + 1
	//
	//fmt.Println("猜猜看我想了哪个数字（1-100）？")
	//
	//// 循环读取用户输入的数字，直到猜中为止
	//for {
	//	var guess int
	//	fmt.Scanln(&guess)
	//
	//	if guess < number {
	//		fmt.Println("猜小了，再试试")
	//	} else if guess > number {
	//		fmt.Println("猜大了，再试试")
	//	} else {
	//		fmt.Println("恭喜你，猜对了！")
	//		break
	//	}
	//}

	// 五言绝句
	//fmt.Println(makePoem())

	// 在常量组中, 如果上一行常量有初始值,但是下一行没有初始值, 那么下一行的值就是上一行的值
	//const (
	//	num1 = 998
	//	num2 // 和上一行的值一样
	//	num3 = 666
	//	num4 // 和上一行的值一样
	//	num5 // 和上一行的值一样
	//)
	//fmt.Println("num1 = ", num1) // 998
	//fmt.Println("num2 = ", num2) // 998
	//fmt.Println("num3 = ", num3) // 666
	//fmt.Println("num4 = ", num4) // 666
	//fmt.Println("num5 = ", num5) // 666

	//const (
	//	num1, num2 = 100, 200
	//	num3, num4 // 和上一行的值一样, 注意变量个数必须也和上一行一样
	//)
	//fmt.Println("num1 = ", num1) //100
	//fmt.Println("num2 = ", num2) //200
	//fmt.Println("num3 = ", num3) //100
	//fmt.Println("num4 = ", num4) //200

	// 利用iota标识符标识符实现从0开始递增的枚举
	//const (
	//	male   = iota // 0
	//	female = iota // 1
	//	yao    = iota // 2
	//)
	//fmt.Println("male = ", male)
	//fmt.Println("female = ", female)
	//fmt.Println("yao = ", yao)

	// 在同一个常量组中,只要上一行出现了iota,那么后续行就会自动递增
	//const (
	//	male   = iota // 这里出现了iota
	//	female        // 这里会自动递增
	//	yao
	//)
	//fmt.Println("male = ", male)     // 0
	//fmt.Println("female = ", female) // 1
	//fmt.Println("yao = ", yao)       // 2

	// 在同一个常量组中,如果iota被中断, 那么必须显示恢复
	//const (
	//	male   = iota
	//	female = 666  // 这里被中断
	//	yao    = iota // 这里显示恢复, 会从当前常量组第一次出现iota的地方开始,每一行递增1, 当前是第3行,所以值就是2
	//)
	//fmt.Println("male = ", male)   // 0
	//fmt.Println("male = ", female) // 666
	//fmt.Println("male = ", yao)    // 2

	//また、iotaは連続した値を生成するためにパターンで使用されることもあります。以下は、パターンでのiotaの使用例です。
	//const (
	//	d = 1 << iota // d = 1 (1 << 0)
	//	e = 1 << iota // e = 2 (1 << 1)
	//	f = 1 << iota // f = 4 (1 << 2)
	//)

	// iota也支持常量组+多重赋值, 在同一行的iota值相同
	//const (
	//	a, b = iota, iota
	//	c, d = iota, iota
	//)
	//fmt.Println("a = ", a) // 0
	//fmt.Println("b = ", b) // 0
	//fmt.Println("c = ", c) // 1
	//fmt.Println("d = ", d) // 1

	//iota自增默认数据类型为int类型, 也可以显示指定类型
	//const (
	//	male float32 = iota // 显示指定类型,后续自增都会按照指定类型自增
	//	female
	//	yao
	//)
	//fmt.Printf("%f\n", male)                       // 0.000000
	//fmt.Printf("%f\n", female)                     // 1.000000
	//fmt.Printf("%f\n", yao)                        // 2.000000
	//fmt.Println("male = ", reflect.TypeOf(female)) // float32

	// 除了和C语言一样,可以通过%o、%x输出八进制和十六进制外, 还可以 直接通过%b输出二进制
	//num := 15
	//fmt.Printf("十进制 = %d\n", num)  // 15
	//fmt.Printf("八进制 = %o\n", num)  // 17
	//fmt.Printf("十六进制 = %x\n", num) // f
	//fmt.Printf("二进制 = %b\n", num)  // 1111

	// Go语言还增加了%T控制符,   用于输出值的类型
	//type Person struct {
	//	name string
	//	age  int
	//}
	//num1 := 10
	//num2 := 3.14
	//per := Person{"lnj", 33}
	//fmt.Printf("num1 = %T\n", num1) // int
	//fmt.Printf("num2 = %T\n", num2) // float64
	//fmt.Printf("per = %T\n", per)   // main.Person

	// Go语言中输出某一个值,很少使用%d%f等, 一般都使用%v即可
	//// 此时相当于把%v当做%d
	//fmt.Printf("num1 = %v\n", num1) // 10
	//// 此时相当于把%v当做%f
	//fmt.Printf("num2 = %v\n", num2) // 3.14

	//num1 := 10
	//num2 := 3.14
	//fmt.Print(num1, num2)                     // 10 3.14
	//fmt.Print("num1 =", num1, "num2 =", num2) // num1 =10 num2 =3.14
	//
	//type Person struct {
	//	name string
	//	age  int
	//}
	//per := Person{"lnj", 33}
	//fmt.Print(per) // {lnj 33}

	// 以下三个函数和Printf/Println/Print函数一样, 只不过上面三个函数是输出到标准输出, 而下面三个函数可以通过w指定输出到什么地方
	//// os.Stdout 写入到标准输出
	//name := "lnj"
	//age := 33
	//// 第一个参数: 指定输出到什么地方
	//// 第二个参数: 指定格式控制字符串
	//// 第三个参数: 指定要输出的数据
	//fmt.Fprintf(os.Stdout, "name = %s, age = %d\n", name, age)
	//
	//// http.ResponseWriter 写入到网络响应
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "name = %s, age = %d\n", name, age)
	//})
	//http.ListenAndServe(":8888", nil)

	// func Scanf(format string, a ...interface{}) (n int, err error)
	//var num1 int
	//var num2 int
	//fmt.Scanf("%d%d", &num1, &num2)
	//fmt.Println(num1, num2)

	// 以下三个函数和Scan/Scanln/Scanf函数一样, 只不过上面三个函数是从标准输入读取数据, 而下面三个函数是从字符串中读取数据
	//str := "lnj 33"
	//var name string
	//var age int
	////fmt.Sscanf(str, "%s %d", &name, &age)
	////fmt.Sscanln(str, &name, &age)
	//fmt.Sscan(str, &name, &age)
	//fmt.Println("name =", name, "age =", age)

	/*
		flag.Xxxx(name, value, usage)
		第一个参数: 命令行参数名称
		第二个参数: 命令行参数对应的默认值
		第三个参数: 命令行参数对应的说明
	*/
	// 1.设置命令行参数
	//name := flag.String("name", "lnj", "请输入人的姓名")
	//age := flag.Int("age", 33, "请输入人的年龄")
	//var name string
	//var age int
	//flag.StringVar(&name, "name", "lnj", "请输入人的姓名")
	//flag.IntVar(&age, "age", 33, "请输入人的年龄")
	//// 2.将命令行参数解析到注册的参数
	//flag.Parse()
	//// 3.使用命令行参数
	//// 注意flag对应方法返回的都是指针类型, 所以使用时必须通过指针访问
	////fmt.Println("name = ", *name) // name =  lnj
	////fmt.Println("age = ", *age)   // age =  33 //注意flag对应方法返回的都是指针类型, 所以使用时必须通过指针访问
	//fmt.Println("name = ", name) // name =  lnj
	//fmt.Println("age = ", age)   // age =  33

	// 新增一个&^运算符
	///*
	//		   0110      a
	//		&^ 1011      b   如果b位位1,那么结果为0, 否则结果为a位对应的值
	//		---------
	//	       0100
	//*/
	//a1 := 6
	//b1 := 11
	//res1 := a1 &^ b1
	//fmt.Println("res1 = ", res1) // 4
	//
	///*
	//		   1011      a
	//		&^ 1101      b   如果b位位1,那么结果为0, 否则结果为a位对应的值
	//		---------
	//	       0010
	//*/
	//a2 := 11
	//b2 := 13
	//res2 := a2 &^ b2
	//fmt.Println("res2 = ", res2) // 2
	// 新增一个&^=运算符   C &^= 2 等于 C = C &^ 2

	//运算符	描述	实例
	//&	返回变量存储地址	&a; 将给出变量的实际地址
	//*	访问指针指向内存	*p; 访问指针p指向内存
	//var num int = 666
	//var p *int = &num
	//fmt.Println(num)
	//fmt.Println(*p)
	//num = 777
	//fmt.Println(num)
	//*p = 999
	//fmt.Println(num)

	// case后面不用编写break, 不会出现case穿透问题
	// 如果想让case穿透,必须在case语句块最后添加fallthrough关键
	//switch num := 3; num {
	//case 1:
	//	fallthrough
	//case 2:
	//	fallthrough
	//case 3:
	//	fallthrough
	//case 4:
	//	fallthrough
	//case 5:
	//	fmt.Println("工作日")
	//case 6:
	//	fallthrough
	//case 7:
	//	fmt.Println("非工作日")
	//default:
	//	fmt.Println("Other...")
	//}

	// case后面不仅仅可以放常量,还可以放变量和表达式
	//value := 2
	//switch num := 1; num {
	//case value: // 变量
	//	fmt.Println("num等于value")
	//default:
	//	fmt.Println("num不等于value")
	//}

	//switch num := 1; num {
	//case getValue(): // 函数①
	//	fmt.Println("num等于value")
	//default:
	//	fmt.Println("num不等于value")
	//}

	//switch num := 18; {
	//case num >= 0 && num <= 10: // 表达式
	//	fmt.Println("num是一个0~10之间的数")
	//case num > 10 && num <= 20:
	//	fmt.Println("num是一个11~20之间的数")
	//default:
	//	fmt.Println("num是一个大于20的数")
	//}

	// case后面定义变量不用添加{}明确作用于范围
	//switch num := 1; num {
	//case 1:
	//	value := 10 // 不会报错
	//	fmt.Println(value)
	//	fmt.Println(num)
	//	value = 120 // 不会报错
	//	fmt.Println(value)
	//default:
	//	fmt.Println("Other...")
	//}

	//// 配列を定義する
	//arr := [8]int{1, 3, 5, 7, 8, 10, 18, 12}
	//// i は、現在のトラバーサルのインデックスを配列に保存するために使用されます
	//// v は、現在トラバースされている値を配列に保存するために使用されます
	//for i, v := range arr {
	//	fmt.Println(i, v)
	//}

	// Go语言中 值类型 有: int系列、float系列、bool、string、数组、结构体
	// 值类型通常在栈中分配存储空间
	// 值类型作为函数参数传递, 是拷贝传递
	// 在函数体内修改值类型参数, 不会影响到函数外的值②
	//arr := [3]int{1, 3, 5}
	//change(arr)
	//fmt.Println(arr) // 1, 3, 5

	// Go语言中 引用类型 有: 指针、slice、map、channel
	// 引用类型通常在堆中分配存储空间
	// 引用类型作为函数参数传递,是引用传递
	// 在函数体内修改引用类型参数,会影响到函数外的值③
	//arr := []int{1, 3, 5}
	//change(arr)
	//fmt.Println(arr) // 1, 8, 5

	// 在函数体内修改引用类型参数,会影响到函数外的值④
	//mp := map[string]string{"name": "lnj", "age": "33"}
	//change(mp)
	//fmt.Println(mp["name"]) // zs

	// 一般的な場合、私たちはほとんどグローバルな匿名関数を使用せず、ほとんどの場合、ローカルな匿名関数を使用します。
	// 匿名関数は、直接呼び出す、変数に保存する、引数として使用する、または戻り値として返すことができます。
	// 変数に保存する
	//a := func(s string) {
	//	fmt.Println(s)
	//}
	//a("hello 変数")
	// 引数として使用する⑤
	//test(func(s string) {
	//	fmt.Println(s)
	//})

	// 戻り値として返す⑥
	//res := test()
	//res(10, 20)

	// 只要闭包还在使用外界的变量, 那么外界的变量就会一直存在⑦
	//res := addUpper() // 执行addUpper函数,得到一个闭包
	//fmt.Println(res())
	//fmt.Println(res())
	//fmt.Println(res())
	//fmt.Println(res())

	//defer fmt.Println("我是第一个被注册的")   // 3
	//fmt.Println("main函数中调用的Println") // 1
	//defer fmt.Println("我是第二个被注册的")   // 2

	// 切片源码
	//type slice struct{
	//	array unsafe.Pointer // 指向底层数组指针
	//	len int // 切片长度(保存了多少个元素)
	//	cap int // 切片容量(可以保存多少个元素)
	//}

	//var arr = [5]int{1, 3, 5, 7, 9}
	//// 从数组0下标开始取,一直取到2下标前面一个索引
	//var sce = arr[0:2]
	//fmt.Println(sce) // [1 3]
	//// 切片len = 结束位置 - 开始位置
	//fmt.Println(len(sce)) // 2 - 0 = 2
	//fmt.Println(cap(sce)) // 5 - 0 = 5
	//// 数组地址就是数组首元素的地址
	//fmt.Printf("%p\n", &arr)    // 0xc00000e450
	//fmt.Printf("%p\n", &arr[0]) // 0xc00000e450
	//// 切片地址就是数组中指定的开始元素的地址
	//// arr[0:2]开始地址为0, 所以就是arr[0]的地址
	//fmt.Printf("%p\n", sce) // 0xc00000e450

	//var arr = [5]int{1, 3, 5, 7, 9}
	//// 同时指定开始位置和结束位置
	//var sce1 = arr[0:2]
	//fmt.Println(sce1) // [1 3]
	//
	//// 只指定结束位置
	//var sce3 = arr[:2]
	//fmt.Println(sce3) // [1 3]
	//
	//// 只指定开始位置
	//var sce2 = arr[0:]
	//fmt.Println(sce2) // [1 3 5 7 9]
	//
	//// 都不指定
	//var sce4 = arr[:]
	//fmt.Println(sce4) // [1 3 5 7 9]

	//方式二: 通过make函数创建make(类型, 长度, 容量)
	//	内部会先创建一个数组, 然后让切片指向数组
	//	如果没有指定容量,那么容量和长度一样
	// 第一个参数: 指定切片数据类型
	// 第二个参数: 指定切片的长度
	// 第三个参数: 指定切片的容量
	//var sce = make([]int, 3, 5)
	//fmt.Println(sce)      // [0 0 0]
	//fmt.Println(len(sce)) // 3
	//fmt.Println(cap(sce)) // 5
	/*
	   内部实现原理
	   var arr = [5]int{0, 0, 0}
	   var sce = arr[0:3]
	*/

	//方式三:通过Go提供的语法糖快速创建　　シンタックスシュガー
	//	和创建数组一模一样, 但是不能指定长度
	//	通过该方式创建时, 切片的长度和容量相等
	//var sce = []int{1, 3, 5}
	//fmt.Println(sce)      // [1 3 5]
	//fmt.Println(len(sce)) // 3
	//fmt.Println(cap(sce)) // 3
	//
	//fmt.Println("追加データ前：", sce)      //  [1 3 5]
	//fmt.Println("追加データ前：", len(sce)) // 3
	//fmt.Println("追加データ前：", cap(sce)) // 3
	//fmt.Printf("追加データ前：%p\n", sce)   // 0xc000010120
	//// 第一个参数: 需要把数据追加到哪个切片中
	//// 第二个参数: 需要追加的数据, 可以是一个或多个
	//sce = append(sce, 666)
	//fmt.Println("追加データ前：", sce)      //  [1 3 5 666]
	//fmt.Println("追加データ前：", len(sce)) // 4
	//fmt.Println("追加データ前：", cap(sce)) // 6   append函数每次给切片扩容都会按照原有切片容量*2的方式扩容
	//fmt.Printf("追加データ前：%p\n", sce)   // 0xc00000e4b0

	// 格式: copy(目标切片, 源切片) , 会将源切片中数据拷贝到目标切片中
	// copy函数在拷贝数据时永远以小容量为准

	//可以通过切片再次生成新的切片, 两个切片底层指向同一数组
	//arr := [5]int{1, 3, 5, 7, 9}
	//sce1 := arr[0:4]
	//sce2 := sce1[0:3]
	//fmt.Println(sce1) // [1 3 5 7]
	//fmt.Println(sce2) // [1 3 5]
	//// 由于底层指向同一数组, 所以修改sce2会影响sce1
	//sce2[1] = 666
	//fmt.Println(sce1) // [1 666 5 7]
	//fmt.Println(sce2) // [1 666 5]

	//var arr []int
	//arr[0] = 2 // panic: runtime error: index out of range [0] with length 0
	//arr[1] = 4
	//arr[2] = 6
	//fmt.Println(arr) // [2 4 6]

	//字符串的底层是[]byte数组, 所以字符也支持切片相关操作
	//str := "abcdefg"
	//// 通过字符串生成切片
	//sce1 := str[3:]
	//fmt.Println(sce1) // defg
	//
	//sce2 := make([]byte, 10)
	//// 第二个参数只能是slice或者是数组
	//// 将字符串拷贝到切片中
	//copy(sce2, str)
	//fmt.Println(sce2) //[97 98 99 100 101 102 103 0 0 0]

}

// 只要闭包还在使用外界的变量, 那么外界的变量就会一直存在⑦
//func addUpper() func() int {
//	x := 1
//	return func() int {
//		x++ // 匿名函数中用到了addUpper中的x,所以这是一个闭包
//		return x
//	}
//}

// 戻り値として返す⑥
//func test() func(int, int) {
//	return func(a int, b int) {
//		fmt.Println(a + b)
//	}
//}

// 引数として使用する⑤
//func test(f func(s string)) {
//	f("hello go")
//}

// 在函数体内修改值类型参数, 不会影响到函数外的值②
//func change(arr [3]int) {
//	arr[1] = 8
//}

// 在函数体内修改引用类型参数,会影响到函数外的值③
//func change(arr []int) {
//	arr[1] = 8
//}

// 在函数体内修改引用类型参数,会影响到函数外的值④
//func change(mp map[string]string) {
//	mp["name"] = "zs"
//}

//①
//func getValue() int {
//	return 1
//}

// 五言绝句
//// 定义一个词库
//var words = []string{"春", "夏", "秋", "冬", "日", "月", "星", "风", "雨", "雪",
//	"山", "水", "云", "花", "草", "鸟", "虫", "鱼", "龙", "虎",
//	"人", "仙", "神", "佛", "道", "情", "意", "心", "梦", "醉",
//	"笑", "泪", "歌", "舞", "琴", "棋", "书", "画", "茶", "酒"}
//
//// 定义一个随机数生成器
//var r = rand.New(rand.NewSource(time.Now().UnixNano()))
//
//// 定义一个函数，用来从词库中随机选择一个词
//func pickWord() string {
//	return words[r.Intn(len(words))]
//}
//
//// 定义一个函数，用来生成一句五言诗
//func makeLine() string {
//	return pickWord() + pickWord() + pickWord() + pickWord() + pickWord()
//}
//
//// 定义一个函数，用来生成一首五言绝句
//func makePoem() string {
//	return makeLine() + "\n" + makeLine() + "\n" + makeLine() + "\n" + makeLine()
//}

package demo

import "fmt"

type Position struct {
	x float64
	y float64
}

func GetMap() {

	var a map[string]int
	fmt.Println(a)                    //输出 map[]
	fmt.Printf("a的长度为:%d \n", len(a)) //默认长度为0
	//cap(a) 错误，字典无法使用cap 函数

	if a == nil { //仅仅声明其 与nil 相等
		fmt.Printf("a与nil 比较\n")
	}
	//a['key1']=1编译错误

	//使用复合字面值初始化 map 类型变量
	a = map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	//初始化后 才可以进行操作
	a["key4"] = 4
	fmt.Println(a)

	//其中使用一些语法糖

	m1 := map[Position]string{
		Position{29.935523, 52.568915}:  "school",
		Position{25.352594, 113.304361}: "shopping-mall",
		Position{73.224455, 111.804306}: "hospital",
	}

	m2 := map[Position]string{
		{29.935523, 52.568915}:  "school",
		{25.352594, 113.304361}: "shopping-mall",
		{73.224455, 111.804306}: "hospital",
	}

	fmt.Println(m1)
	fmt.Println(m2)

}

func GetMapByMake() {
	var b = make(map[string]int)
	fmt.Println(b)                    //输出 map[]
	fmt.Printf("b的长度为:%d \n", len(b)) //默认长度为0
	if b == nil {                     //初始化后,与nil不等
		fmt.Printf("b与nil 比较")
	}

	c := make(map[string]int, 100)    //声明后指定容量
	fmt.Println(c)                    //输出 map[]
	fmt.Printf("c的长度为:%d \n", len(c)) //默认长度为0
	if c == nil {                     //初始化后,与nil不等
		fmt.Printf("c与nil 比较")
	}

	//插入新键值对

	b["key1"] = 1
	b["key2"] = 2
	b["key3"] = 3
	fmt.Println(b)

}

var exampleMap = map[string]int{
	"key1": 1,
	"key2": 2,
	"key3": 3,
}

//关于map的操作
func MapOperations() {
	//新增
	exampleMap["key4"] = 4
	fmt.Println(exampleMap) //map[key1:1 key2:2 key3:3 key4:4]

	//修改
	exampleMap["key1"] = 5
	fmt.Println(exampleMap) //map[key1:5 key2:2 key3:3 key4:4]

	//删除
	delete(exampleMap, "key2") // 删除"key2"
	fmt.Println(exampleMap)    //map[key1:5 key3:3 key4:4]
	//delete 函数是从 map 中删除键的唯一方法。即便传给 delete 的键在 map 中并不存在，delete 函数的执行也不会失败，更不会抛出运行时的异常
	delete(exampleMap, "key10") // 删除"key10"
	fmt.Println(exampleMap)     //map[key1:5 key3:3 key4:4]

	//查找

	v := exampleMap["key10"]
	fmt.Printf("key10的value为:%d \n", v)
	//如果这个键在 map 中并不存在，我们也会得到一个值，这个值是 value 元素类型的零值。

	//应该使用下面的方法
	value, ok := exampleMap["key1"]
	if ok { // 找到了
		fmt.Printf("key1的value为:%d \n", value)
	}

	//遍历map
	//同一 map 做多次遍历的时候，每次遍历元素的次序都不相同
	//程序逻辑千万不要依赖遍历 map 所得到的的元素次序
	for key, value := range exampleMap {
		fmt.Println(key, value)
	}

	//借助匿名变量只获取字典的值
	for _, value := range exampleMap {
		fmt.Println(value)
	}
	//只获取字典的键名
	for key := range exampleMap {
		fmt.Println(key)
	}

	//注意点
	//map 不是线程安全的，不支持并发读写；
	//不要尝试获取 map 中元素（value）的地址
	//p := &exampleMap["key1"] // cannot take the address of m[key]fmt.Println(p)
	//fmt.Println(p)

}

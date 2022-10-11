package object

import (
	"fmt"
	"time"
	"unsafe"
)

type User struct {
	name     string //字段的首字母的大小写来判定这个标识符是否为导出标识符
	password string
	mail     string
}

//实例化结构体
func GetStruct() {
	//方法1  var 声明
	var user User
	//与数组类型相同，结构体类型属于值类型，因此结构体类型的零值不是 nil，上述  User 的零值就是  User{}
	fmt.Println(user)
	user.name = "姓名"
	user.password = "密码"
	user.mail = "邮箱"
	fmt.Println(user)

	//方法2. new // 返回指针型变量
	user1 := new(User)
	user1.name = "姓名1"
	user1.password = "密码1"
	user1.mail = "邮箱1"
	fmt.Println(user1)

	//方法三  赋值初始化
	user2 := User{"姓名2", "密码2", "邮箱2"}
	fmt.Println(user2)

	user3 := User{name: "姓名3", password: "密码3", mail: "邮箱3"}
	fmt.Println(user3)

	// 方法三中第一种方式
	//必须初始化结构体的所有字段。
	//每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	//	user4:=User{"姓名4","密码4"}

	//方法4 &   //返回指针型变量
	var user5 = &User{}
	user5.name = "姓名5"
	user5.password = "密码5"
	fmt.Println(user5)

}

//空结构体
type Empty struct{} // Empty是一个不包含任何字段的空结构体类型

func GetEmptyStruct() {

	var s Empty
	println(unsafe.Sizeof(s)) // 0
	//空结构体类型变量的内存占用为 0。
	//基于空结构体类型内存零开销这样的特性，
	//我们在日常 Go 开发中会经常使用空结构体类型元素，作为一种“事件”信息进行 Goroutine 之间的通信

	var c = make(chan Empty) // 声明一个元素类型为Empty的channel
	c <- Empty{}             // 向channel写入一个“事件”
}

//匿名结构体
//如果只是临时使用struct一次，而不是多次使用，用匿名struct即可

func GetAnonymousStruct() {
	user := struct {
	}{}
	fmt.Println(user)
	user1 := struct {
		name string
		age  int
	}{
		"姓名",
		12,
	}
	fmt.Println(user1)
}

type T struct {
	b byte
	i int64
	u uint16
}

type S struct {
	b byte
	u uint16
	i int64
}

//Go 语言中结构体类型的大小受内存对齐约束的影响
func GetSizeStruct() {
	var t T
	println(unsafe.Sizeof(t)) // 24
	var s S
	println(unsafe.Sizeof(s)) // 16
}

//结构体标签
// json(JSON标签) orm(Beego标签)、gorm(GORM标签)、bson(MongoDB标签)、form(表单标签)、binding(表单验证标签)

//普通标签
type Teacher struct {
	ID   int    "工号"
	Name string "姓名"
	Age  int
}

//JSON 标签
type Student struct {
	ID   int    `json:"-"`           // 该字段不进行序列化
	Name string `json:name,omitempy` // 如果为类型零值或空值，序列化时忽略该字段
	Age  int    `json:age,string`    // 指定类型，支持string、number、boolen
}

//gorm标签
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//form标签
type LoginForm struct {
	Email    string `form:"emial"`
	Password string `form:"password"`
}

//binding
type People struct {
	Name string `form:"name" binding:"required,min=1,max=10"`
	Age  int    `form:"age" binding:"lte=150,gte=0"`
	sex  string `form:"sex" binding:"oneof=male female"`
}

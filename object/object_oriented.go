package object

import "fmt"

//使用嵌入字段实现面向对象 这里其实本质是go 组合的思想

type Animal struct {
	Name string
}
func (a Animal) Call() string {
	return "动物的叫声..."
}
func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}
func (a Animal) GetName() string  {
	return a.Name
}

type Dog struct {
	Animal
}

func GetInheritance() {
	animal := Animal{"中华田园犬"}
	dog := Dog{animal}
	fmt.Println(dog.GetName())
	fmt.Println(dog.Call())
	fmt.Println(dog.FavorFood())
}

type Dog1 struct {
	Animal
}


func (d Dog1) FavorFood() string {
	return "骨头"
}
func (d Dog1) Call() string {
	return "汪汪汪"
}

func GetPolymorphic() {
	animal := Animal{"中华田园犬"}
	dog := Dog1{animal}
	fmt.Println(dog.GetName())
	fmt.Println(dog.Call())
	fmt.Println(dog.FavorFood())

}
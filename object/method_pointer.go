package object

type T1 struct {
	a int
}

func (t T1) M1() {
	t.a = 10
}

func (t *T1) M2() {
	t.a = 11
}

//receiver 参数类型对 Go 方法的影响
//如果 Go 方法要把对 receiver 参数代表的类型实例的修改，反映到原类型实例上，那么我们应该选择 *T 作为 receiver 参数的类型。
//如果 receiver 参数类型的 size 较大，以值拷贝形式传入就会导致较大的性能开销，这时我们选择 *T 作为 receiver 类型可能更好些。

func GetMethodPointer() {
	var t T1
	println(t.a) // 0

	//当 receiver 参数的类型为 T 时
	//当我们的方法 M1 采用类型为 T 的 receiver 参数时，
	//代表 T 类型实例的 receiver 参数以值传递方式传递到 M1 方法体中的，
	//实际上是 T 类型实例的副本，M1 方法体中对副本的任何修改操作，都不会影响到原 T 类型实例
	t.M1()
	println(t.a) // 0

	//第二，当 receiver 参数的类型为 *T 时
	//当我们选择以 *T 作为 receiver 参数类型时，
	//M2 方法等价转换为F2(t *T)。
	//都会反映到原 T 类型实例上
	p := &t
	p.M2()
	println(t.a) // 11




}


type T2 struct {
	a int
}

func (t T2) M1() {
	t.a = 10
}

func (t *T2) M2() {
	t.a = 11
}


//无论是 T 类型实例，还是 *T 类型实例，都既可以调用 receiver 为 T 类型的方法，也可以调用 receiver 为 *T 类型的方法
func GetMethodPointerV1() {


	//T 类型的实例 t1 之所以可以调用 receiver 参数类型为 *T 的方法 M2，
	//都是 Go 编译器在背后自动进行转换的结果。或者说，t1.M2() 这种用法是 Go 提供的“语法糖”：
	//Go 判断 t1 的类型为 T，也就是与方法 M2 的 receiver 参数类型 *T 不一致后，会自动将t1.M2()转换为(&t1).M2()。
	var t1 T2
	println(t1.a) // 0
	t1.M1()
	println(t1.a) // 0
	t1.M2()
	println(t1.a) // 11

	var t2 = &T2{}
	println(t2.a) // 0
	t2.M1()
	println(t2.a) // 0
	t2.M2()
	println(t2.a) // 11
}

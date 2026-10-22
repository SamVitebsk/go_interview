package main

type Closer interface {
	Close()
}

type File struct{}

func (f *File) Close() {}

type Foo struct{}

func (*Foo) A() {}
func (*Foo) B() {}
func (*Foo) C() {}

type AB interface {
	A()
	B()
}

type BC interface {
	B()
	C()
}

func main() {
	var f AB = &Foo{}
	y := f.(BC)
	_ = y
	// y.A()

	// 1. Сработает ли type assertion f.(BC)?
	//    ДА. В runtime f содержит пару (dynamic_type: *Foo, dynamic_value: &Foo{}).
	//    Набор методов *Foo содержит A(), B(), C(). Интерфейс BC требует B() и C().
	//    Так как *Foo удовлетворяет BC, type assertion f.(BC) завершается успешно.
	//
	// 2. Доступен ли вызов y.A()?
	//    НЕТ. Это ошибка компиляции: "y.A undefined (type BC has no field or method A)".
	//    Компилятор проверяет доступность методов строго по СТАТИЧЕСКОМУ типу переменной y (интерфейс BC),
	//    а не по динамическому типу лежащего в ней значения (*Foo).

}

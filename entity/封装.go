package entity

//可以模仿Java中的封装,但又比Java的封装更严格，只对外暴露的是包，包里的结构体类是不直接暴露的

type order struct {
	orderName   string //私有的，相当于private
	orderNumber string
}

func (o *order) GetOrderName() string {
	return o.orderName
}
func (o *order) GetOrderNumber() string {
	return o.orderNumber
}
func (o *order) SetOrderNumber(number string) {
	o.orderNumber = number
}
func (o *order) SetOrderName(name string) {
	o.orderName = name
}
func NewOrder(name string) *order { //相当于是一个构造函数
	return &order{orderName: name}
}

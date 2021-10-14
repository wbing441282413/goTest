package interface_entity

import "fmt"

type AliPay struct { //阿里支付实体

}

type WeChatPay struct { //微信支付实体

}

type AliPayer interface {
	AliPayFunc()
}

type WeChatPayer interface {
	WeChatPayFunc()
}

func (w *AliPay) AliPayFunc() { //实现了阿里支付接口
	fmt.Printf("%T支付\n", w)
}
func (w *WeChatPay) WeChatPayFunc() {
	fmt.Printf("%T支付\n", w)
}

func Pay(inter interface{}) { //使用Switch对接口类型进行断言后调用
	switch inter.(type) { //利用inteface接口类型可以进行类型断言的特点，结合switch来判断调用什么支付接口
	//注意的是，interface.(type)返回的是inteface的接口类型，也就是说结构体实现了接口后，这里返回的是接口类型不是结构体
	case AliPayer:
		pay, ok := inter.(AliPay) //可以看到接口的接收者是指针，那他在断言的时候会被认为是指针
		/**
		试想一下，如果接口绑定接受者是结构体值，那这里结果还是ok=false
		因为，在main中调用的时候传入的是指针，由于接口接受者是值的时候，不论是指针还是值都可以赋值给接口，go会判断并转换
		所以，由于传入的是指针，所以指针是被指针赋值的，所以这里还是指针，所以这里还是false
		*/
		var p *AliPay
		if !ok {
			fmt.Println("断言转换错误！！！！！panic ,试着转换成*AliPay")
			p = inter.(*AliPay)
			p.AliPayFunc()
		} else {
			pay.AliPayFunc()
		}
	case WeChatPayer:
		pay := inter.(*WeChatPay)
		pay.WeChatPayFunc()
	}
}

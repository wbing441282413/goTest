package entity

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

func Pay(inter interface{}) { //Switch对接口类型进行断言后调用
	switch inter.(type) {
	case AliPayer:
		pay, ok := inter.(AliPay)
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

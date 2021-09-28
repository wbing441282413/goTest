package entity

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

type ReaderAndClose interface { //组合接口
	Reader
	Closer //接口内嵌
}

package main

import (
	"bufio"
	"container/list"
	"encoding/json"
	"fmt"
	"go/types"
	"gotest/entity"
	"gotest/utils"
	"io"
	"log"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Stu struct {
	id   int
	name string
}

func Test(s Stu) {
	s.id = 1
	s.name = "yy"
}

func Test2(s *Stu) {
	s.id = 2
	s.name = "yang"
}
func main() {
	student := new(entity.Student)
	student.Name = "王兵"
	fmt.Println(student.Name)

	mp := make(map[string]int)
	mp["A"] = 1
	mp["b"] = 2
	delete(mp, "A")
	for name, age := range mp {
		fmt.Println(name)
		fmt.Println(age)
	}

	entity.Say()

	fmt.Println("student:	", entity.GetName(*student))

	student2 := entity.Student{1, "sdsd"}
	entity.Swap(&student2, student) //与结构体相关的不能直接调用
	fmt.Println("student:	", entity.GetName(*student))

	var age int = 10
	var p *int = &age
	*p++ //对指针操作直接是对值操作
	x := &p
	fmt.Println(x)
	fmt.Println(age)

	//var s Stu
	s := new(Stu)
	pp := Stu{1, "wb"}
	var t Stu
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(reflect.TypeOf(pp))
	var pe = &Stu{}
	fmt.Println(reflect.TypeOf(pe))

	song := entity.Song{"哈哈哈"}
	singer := entity.Singer{3, song}
	fmt.Println(song.GetSongName())
	singer.Intro()

	var sing entity.Sing
	sing = singer
	sing.Sing()

	//不论是定义在哪里的接口，只要绑定了就能用
	var tt entity.Tt
	tt = singer
	tt.Tt()
	tt = song
	tt.Tt()

	if 1 == 1 && 1 == 1 {
		fmt.Println("ccccccc")
	}

	var score int = 90
	switch {
	case score >= 90:
		fmt.Println("成绩优良")
		fallthrough
	case score > 70 && score < 90:
		fmt.Println("成绩合格") //如果在case后面使用fallthrough 关键字，则当前case执行完成后会继续执行下一个case，
	case score < 60:
		fmt.Println("成绩不及格")
	}

	var aa [3]int = [3]int{1, 2, 3}
	cc := [...]int{1, 2}
	var dd [2]int //默认初始是0
	fmt.Println(reflect.TypeOf(aa))
	fmt.Println(reflect.TypeOf(cc))
	for _, a := range dd {
		fmt.Println(a)
	}

	//数组
	k := [2][2]int{{1, 1}, {2, 2}}
	var j [2]int
	j = k[1]
	for _, a := range j {
		fmt.Println(a)
	}

	var gg [2]int
	for _, a := range gg {
		fmt.Println(a)
	}

	//未初始化的结构体
	var xx entity.Singer
	fmt.Println("没初始化的结构体：", xx.Song.SongName, "count:", xx.Count)

	slice := make([]int, 2, 3)
	fmt.Println("=========切片========")
	for _, a := range slice {
		fmt.Println(a)
	}

	ll := []int{1, 2, 3, 4, 5}
	l := ll[2:5]
	for _, a := range l { //range返回的是切片元素的复制，而不是元素的引用，所以这里我们修改v的值并不会改变slice切片里的值
		fmt.Println(a)
	}

	//切片追加
	fmt.Println("-----------切片追加-----------")
	sl := []int{1, 2, 3, 4, 5}
	newSlice := sl[1:2:3]
	newSlice = append(newSlice, sl...)
	fmt.Println(newSlice)
	fmt.Println(sl)

	for _, a := range sl {
		a = a + 1
	}
	fmt.Println(sl) //可以看到没有变

	//排序切片
	sort.Ints(newSlice)
	fmt.Println(newSlice)

	sx := [][]int{{10}, {100, 200}}
	sx[0] = append(sx[0], 20)
	fmt.Println(sx)

	//待查询的数据：
	//数据内容不重复
	lis := []entity.Profile{
		{Name: "张三", Age: 23, Address: "cq"},
		{Name: "李四", Age: 25},
		{Name: "王五"},
	}
	//传统查询
	entity.FindData(lis, "张三", 23)
	//利用map的多键索引查询（组合键查询）
	entity.BuildIndex(lis)     //构建基于查询的组合键（name、age)
	entity.QueryData("张三", 23) //依据name、age进行查询（多条件）

	L1 := list.New()
	var L2 list.List
	fmt.Println(reflect.TypeOf(L1))
	fmt.Println(reflect.TypeOf(L2))
	fmt.Println(L1)
	fmt.Println(L2)

	//循环
	//for i:=1 ; i<=9;i++{
	//	for j:=1 ; j<=i;j++{
	//		fmt.Printf("%dX%d=%d	",j,i,j*i)
	//	}
	//	fmt.Println()
	//}

	//ss:= "sdsd"
	//for _ ,e:= range ss{
	//	fmt.Printf("%c ",e)
	//}
	//zx := make(map[string]int)
	//zx["w"] = 1
	//zx["c"] = 2
	//for k,v := range zx{
	//	fmt.Println(k,v)
	//}
	//Switch
	//var ax = 1
	//switch ax{
	//	case 1,2:
	//		fmt.Println("xxxxxx")
	//	case 3:
	//		fmt.Println("ccccc")
	//}
	shu := []int{1, 2} //切片是指针，所以修改能成功
	swap(shu)
	fmt.Println(shu)

	shu2 := [2]int{1, 2} //数组是值修改不会成功
	swap2(shu2)
	fmt.Println(shu2)

	ii := [3]int{1, 2, 3}
	aaa := ii[:] //运行结果看到切片并不会开辟新的内存空间
	fmt.Println(reflect.TypeOf(ii))
	fmt.Println(reflect.TypeOf(aaa))
	fmt.Printf("%p \n", &ii)
	fmt.Printf("%p \n", aaa)

	//二分法
	yy := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(utils.BinarySearch(yy[:], 0, 7, 8))
	fmt.Println(utils.BinarySearchRecursion(yy[:], 0, 7, 8))

	//类似于Java中的匿名函数
	utils.ToSay(func() {
		fmt.Println("我在帮别人说话")
	})

	mpp := make(map[interface{}]interface{})
	mpp["sd"] = "Sda"
	mpp[1] = "Sd"
	mpp[2] = 1
	fmt.Println(mpp)
	fmt.Printf("%T", mpp["sd"])
	fmt.Printf("%T", mpp[1])
	fmt.Printf("%T \n", mpp[2])

	vv, kkk := mpp["sd"].(int) //空接口断言，不是int则kkk为false
	if kkk {
		fmt.Println(vv)
	}

	var pxp = new(entity.Phone) //结构指针
	pxp.Name = "小米"
	var us entity.Usber
	us = pxp
	us.Start() //也可以

	m1 := &entity.Phone{Name: "xx"}
	m2 := new(entity.Phone)
	fmt.Println(reflect.TypeOf(m1))
	fmt.Println(reflect.TypeOf(m2))

	months := [...]string{1: "January", 12: "December"} //数组
	//summer := months[3:5]//summer长度是2
	//fmt.Println(summer[3])//切片越界
	fmt.Println(len(months))

	bb := []int{1, 2, 3, 4, 5}
	utils.Reverse(bb[1:4])
	fmt.Println(bb)

	var ggx []int
	if ggx == nil {
		fmt.Println("没初始化的slice默认是nil")
		fmt.Println(len(ggx))
		fmt.Println(cap(ggx))
	}

	//utils.NoDo()

	//append的演示
	var xs []int
	for i := 0; i < 10; i++ {
		xs = append(xs, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(xs), xs)
	}

	so := []string{"sd", "sd", "ss", "ss", "xx"}
	//utils.DeleteMultiChar(so)
	fmt.Println(so)

	sss, hhh := "sd", "sd"
	if sss == hhh {
		fmt.Println("判断字符串相等可以用==")
	}

	fmt.Println(utils.Equal(map[string]int{"A": 0}, map[string]int{"B": 1}))

	var sy string
	sy = fmt.Sprintf("%q", so)
	fmt.Println(sy)

	maa := make(map[int]string)
	maa[1] = "uuu"
	utils.ChangeMap(maa)
	fmt.Println(maa) //可见是指针类型的

	//json,多个结构体实体
	var movies = []entity.Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Printf("%T\n", data)
	//解码
	var movies2 []entity.Movie
	if err = json.Unmarshal(data, &movies2); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(movies2)

	var actors []struct{ Actors []string }
	if err = json.Unmarshal(data, &actors); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(actors)

	var ree []struct{ released string }
	if err = json.Unmarshal(data, &ree); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(ree)

	//uy := new([]entity.Song)
	uy := make([]entity.Song, 10)
	uy[1].SongName = "Sd"
	fmt.Println("-----", reflect.TypeOf(uy))
	fmt.Println("-----", uy)

	//inEOF()
	fmt.Println(runtime.GOARCH)  //CPU型号
	fmt.Println(strconv.IntSize) //int位数

	//标准输入scan
	//var tr string
	//var na int
	//fmt.Scan(&tr, &na)
	//fmt.Println(tr,na)

	//切片的标准输入
	//var buffer [10]byte
	//n, err := os.Stdin.Read(buffer[:])
	//if err != nil {
	//	fmt.Println("read error:", err)
	//	return
	//}
	//fmt.Println("count:", n, ", msg:", string(buffer[:]))

	//通过bufio读取
	//reader := bufio.NewReader(os.Stdin)
	//result, err := reader.ReadString('\n')
	//if err != nil {
	//	fmt.Println("read error:", err)
	//}
	//fmt.Println("result:", result)
	//fmt.Println(reflect.TypeOf(result))

	//strings.Map对字符串中的每个字符调用addd函数，并将每个addd函数的返回值组成一个新的字符串返回给调用者
	fmt.Println(strings.Map(addd, "abc efg"))

	ccd()

	fmt.Println(f2())
	fmt.Println(f3())

	p11 := entity.ColorPoint{&entity.Point{2, 2}, "red"}
	p22 := entity.ColorPoint{&entity.Point{1, 1}, "yellow"}
	fmt.Printf("%p\n", &p11)
	fmt.Printf("%p\n", &p22)
	fmt.Printf("%p\n", p11.Point)
	fmt.Printf("%p\n", p22.Point)
	p11.Point = p22.Point //此时的p11的point指向的是p22的point地址
	fmt.Println("--把成员point赋值给对方后---")
	fmt.Printf("%p\n", &p11)
	fmt.Printf("%p\n", &p22)
	fmt.Printf("%p\n", p11.Point)
	fmt.Printf("%p\n", p22.Point) //可以看到成员的地址变一样了
	//fmt.Println(&p22.Point)//指针不应该加&取址，值才要取址

	//add 和 ADD 的底层类型相同，并且 add 是字面量类型
	//所以 add 可直接赋值给 ADD 类型的变量 g
	var g entity.ADD = entity.Adds
	fmt.Println(g(1, 2))

	//不一定要一个类型全部实现接口的所有方法，可以提过内嵌其他按类型，让其他类型实现的也可以为我所用
	AF := new(entity.Aa)
	var sssd entity.Service = AF
	sssd.Start()
	sssd.Stop()
	var xxxd entity.Habe = AF
	xxxd.Show()

	//自己和内嵌都实现了接口会调用哪个
	gh := new(entity.YY)
	var xxa = gh
	xxa.HH()

	//时间相关
	var currentTime time.Time = time.Now()
	fmt.Println(currentTime)

	var yu interface{}
	if yu == nil {
		fmt.Println("空接口等于nil")
	}

	switch yu.(type) {
	case types.Nil:
		fmt.Println("空接口类型是Nil")
	default:
		fmt.Println("unknown type")
	}
	entity.TimeSomething()

	//fp := "D:\\Backup\\Documents\\草稿.txt"
	//ddda, err := entity.ReadFileByFilePath(fp)//小文件一次性读出
	//fmt.Printf("%s\n",ddda)

	entity.SortUse()

	//fmt.Println(reflect.TypeOf(entity.Yuu))
	//entity.YYY(entity.XXC)//函数签名和函数类型的不一样时是不能转换的
	entity.HJ(entity.YYY(entity.Yuu))

	//Switch结合类型断言
	entity.Pay(new(entity.AliPay))
	entity.Pay(new(entity.WeChatPay))

	//一个类型实现了多个接口的话，那他就可以是多个接口类型的
	animals := map[string]interface{}{
		"bird": new(entity.Bird),
		"dog":  new(entity.Dog),
	}
	for name, obj := range animals {
		flyer, isFlyer := obj.(entity.Fly)
		walker, isWalker := obj.(entity.Walk)
		fmt.Printf("name:	%s,type:	%T,isFlyer:	%v\n", name, flyer, isFlyer)
		fmt.Printf("name:	%s,type:	%T,isWalker:	%v\n", name, walker, isWalker)
	}

	//自定义error接口的实现类来实现自定义异常
	sqrt, err := entity.Sqrt(-13.23)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrt)
	}

	//简易的web,用http包
	//entity.SendToWeb()

	//封装
	//order := entity.NewOrder("篮球")
	//fmt.Println(order.GetOrderName())
	//order.SetOrderName("足球")
	//fmt.Println(order.GetOrderName())

	//因为在程序启动时，Go 程序就会为 main() 函数创建一个默认的 goroutine,程序默认是在main的这个协程中运行的
	//当有go作用时，这个被作用的部分就会新建一个协程，被作用的部分就在这个新建的协程中运行了，然后main的协程和这个新建的协程并发运行
	//go entity.Caculate()
	entity.Caculate()

	//感受一下资源竞争,没有同步措施的
	//wg.Add(2)//计数器为负时会爆出panic
	//go incCount()
	//go incCount()
	//wg.Wait()//Wait() 一直阻塞，直到wg为零——也就是上面两个协程都运行完
	//fmt.Println(count)
	//
	////有同步措施的，原子函数
	//wg.Add(2)
	//go incSycCount()
	//go incSycCount()
	//wg.Wait()
	//fmt.Println(counter)

	wg.Add(2)
	go mutexCount()
	go mutexCount()
	wg.Wait()
	fmt.Println(cou)

	//runtime.GOMAXPROCS(runtime.NumCPU())//多核执行
	fmt.Println(runtime.NumCPU()) //本机CPU核数

	//通道创建
	//ch1 := make(chan interface{})//传输任意类型的数据
	//ch1 <- 0

	//time.Sleep(time.Duration(5) * time.Second)

	//通信
	ch := make(chan interface{})
	//go func() {
	//	fmt.Printf("%d----开始发消息\n",utils.Goid())
	//	ch <- 0//有人接收走通道中的数据之前一直阻塞在这,一旦被接收就继续向下执行
	//	for{
	//		fmt.Printf("%d继续工作--\n",utils.Goid())
	//	}
	//}()
	//time.Sleep(time.Nanosecond)
	//xh := <- ch
	//for i := 0; i < 10; i++ {
	//	fmt.Println("main工作中")
	//}
	//time.Sleep(time.Nanosecond)
	//fmt.Println("MAIN收到消息了：	",xh)

	//通过range来接收消息
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	for data := range ch { //相当于是data <- ch
		fmt.Println(data)
		if data == 3 {
			break
		}
	}
}

var (
	count   int
	counter int32
	cou     int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

//非同步的方式
func incCount() {
	defer wg.Done() //函数每次执行完wg计数器都会减一
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched() //暂停执行，让别的协程运行,如果别的协程运行完了会回到这里继续向下执行
		value++
		count = value
	} //出现问题的主要原因是数据在每个协程中没有同步
}

//传统的同步方式：加锁的方式实现同步
//有2种方式原子函数和互斥锁
func incSycCount() { //原子函数的方式
	defer wg.Done() //函数每次执行完wg计数器都会减一
	for i := 0; i < 2; i++ {
		fmt.Println("当前协程id是：	", utils.Goid())
		atomic.AddInt32(&counter, 1)
		runtime.Gosched() //暂停执行，让别的协程运行,如果别的协程运行完了会回到这里继续向下执行
	}
}

func mutexCount() { //互斥方式，加锁机制
	fmt.Printf("---%d协程开始执行\n", utils.Goid())
	defer wg.Done()
	for i := 0; i < 5; i++ {
		mutex.Lock()
		{
			fmt.Println("Gosched--前--当前协程id：", utils.Goid())
			value := cou
			fmt.Printf("---%d协程退出执行\n", utils.Goid())
			runtime.Gosched() //会退出协程给别的协程操作，但是别的协程没有锁执行不了，所以会回来继续执行这个协程
			fmt.Printf("---%d回到恢复执行\n", utils.Goid())
			value++
			cou = value
			fmt.Println("Gosched--后--当前协程id：", utils.Goid())
			fmt.Println("把cou修改为：	", cou)
		}
		fmt.Printf("---%d协程第%d次迭代结束，释放锁\n", utils.Goid(), i)
		mutex.Unlock() //这里好像释放之后，其他协程也一直获取不到锁，是因为执行的太快了,下面sleep一下就可以有效果了
		time.Sleep(time.Second)
	}
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}() //匿名函数并调用
	return t
}
func f3() (t int) {
	t = 5
	defer func() {
		t = t + 5
	}()
	return t
}

func ccd() {
	i := 1
	defer fmt.Println(i)
	fmt.Println("我先运行完的")
	return
}

func addd(r rune) rune {
	return r + 1
}

func swap(a []int) {
	tmp := a[0]
	a[0] = a[1]
	a[1] = tmp
}

func swap2(a [2]int) {
	tmp := a[0]
	a[0] = a[1]
	a[1] = tmp
}

func inEOF() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return fmt.Errorf("read failed:%v", err)
		}
		// ...use r…
		fmt.Println(r)
	}
	return nil
}

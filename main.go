package main

//--
import (
	"bufio"
	"encoding/json"
	"fmt"
	"gotest/entity"
	"gotest/utils"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime"
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

	//二分法
	yy := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(utils.BinarySearch(yy[:], 0, 7, 8))
	fmt.Println(utils.BinarySearchRecursion(yy[:], 0, 7, 8))

	//utils.NoDo()

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

	//add 和 ADD 的底层类型相同，并且 add 是字面量类型
	//所以 add 可直接赋值给 ADD 类型的变量 g
	var g entity.ADD = entity.Adds
	fmt.Println(g(1, 2))

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

	// utils.RedisOps()

	fmt.Println("模拟一下消费者生产者模式")
	cch := make(chan int)
	go Consumer(cch)
	//假设是生产者
	for i := 1; i < 5; i++ {
		fmt.Printf("当前通道元素的长度：	%d\n", len(cch))
		time.Sleep(time.Second)
		cch <- i
	}

	cch <- 0 //生产者传入0，告知消费者，不再生产了
	<-cch    //这里是必须要消费的,否则消费者传的结束信息不被消费会一直等待

	fmt.Println("到底非缓冲通道中是不是只有一个元素")
	kk := make(chan int)
	go dudu(kk)
	kk <- 0
	// <-kk //会报错，同一个协程不能自己发自己收

	// fmt.Println("通道的只读和只写")
	// jk := make(chan<- int) //<-chan int就是只读的通道
	// jk <- 0
	// <- jk //会提示错误，只写通道不能读
	// fmt.Println(len(jk))

	//无缓冲通道实现打乒乓球
	court := make(chan int) //模拟比赛
	wgg.Add(2)
	go Player(court, "马龙")
	go Player(court, "樊振东")

	court <- 1 //发球，比赛开始
	wgg.Wait()

	//无缓冲通道模拟接力跑
	// 创建一个无缓冲的通道
	baton := make(chan int)
	// 为最后一位跑步者将计数加1
	wgg.Add(1)
	// 第一位跑步者持有接力棒
	go Runner(baton)
	// 开始比赛
	baton <- 1
	// 等待比赛结束
	wgg.Wait()

	fmt.Println("开启多个分支完成计算任务")
	wgg.Add(5)
	he := 0
	cj := make(chan int, 8)
	fmt.Println("还没插入元素的时候，通道cj的长度是：", len(cj)) //0
	for i := 0; i < 10000; {
		if i%2000 == 0 {
			go Cal(i, i+1999, cj)
		}
		i = i + 2000
	}
	wgg.Wait() //等运算协程都算完再加起各个结果
	// for xxum := range cj {
	// 	he += xxum
	// } //感觉这种方式有问题，还是不用为好，这里用的话死锁了

	for len(cj) > 0 {
		a := <-cj
		he += a
	}
	fmt.Println("计算完后，通道cj的长度是：	", len(cj)) //可以看到即便是由缓存的通道，元素读完后出的也是0 了
	//看源码可以知道，len取的是元素的长度，不是缓存区的长度，
	fmt.Println(he)
}
func Consumer(ch chan int) { //假装是消费者
	for { //一直消费
		data := <-ch
		if data == 0 {
			break
		}

	}
	ch <- 0 //告诉生产者我知道了你不生产了，那我告诉你我不消费了
}
func Cal(start int, end int, ch chan int) { //模拟一个计算任务，计算加一10000次
	defer wgg.Done()
	sum := 0
	for i := start; i <= end; i++ {
		sum += 1
	}
	ch <- sum
	fmt.Println("本协程的计算结果是：	", sum)
	fmt.Println("通道元素长度：	", len(ch))
}

func Player(ch chan int, name string) {
	defer wgg.Done()

	for {
		ball, ok := <-ch
		if !ok { //通道被关闭，也就是对手关闭了通道，我们赢了
			fmt.Printf("%s 赢了\n", name)
			return
		}

		i := rand.Intn(100)
		if i%13 == 0 { //模拟失球
			fmt.Printf("%s 输了\n", name)
			close(ch)
			return
		}
		fmt.Printf("%s ---击球\n", name)
		ball++
		ch <- ball //把球打给对手
	}

}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int
	// 等待接力棒
	runner := <-baton
	// 开始绕着跑道跑步
	fmt.Printf("Runner %d Running With Baton\n", runner)
	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton) //创建的新跑者会阻塞在runner:=<-baton，直到当前跑者执行完baton <- newRunner才会向下执行
	}
	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)
	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wgg.Done()
		return
	}
	// 将接力棒交给下一位跑步者
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		runner,
		newRunner)
	baton <- newRunner
}

var (
	wgg sync.WaitGroup
)

func dudu(kk chan int) {
	fmt.Printf("----当前通道元素的长度：	%d\n", len(kk))
	time.Sleep(time.Second)
	<-kk
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

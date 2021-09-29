package entity

/**
并发打印
*/
func Consumer(ch chan int) { //假装是消费者
	for { //一直消费
		data := <-ch
		if data == 0 {
			break
		}
	}
	ch <- 0 //告诉生产者我知道了你不生产了，那我告诉你我不消费了
}

package main

import (
	"github.com/weiwenwang/ProductAndConsumer/Mysql"
	"github.com/weiwenwang/ProductAndConsumer/Redis"
	Commodity3 "github.com/weiwenwang/ProductAndConsumer/Redis/Commodity"
	Commodity2 "github.com/weiwenwang/ProductAndConsumer/Commodity"
	"github.com/weiwenwang/ProductAndConsumer/Mysql/Commodity"
	"fmt"
	"sync"
)

const CONSUMER_NUMBER = 1 // 消费者的数量
const PRODUCT_NUMBER = 20 // 生产者的数量

func main() {
	wg := sync.WaitGroup{}
	Mysql.InitDB()    // 初始化Mysql
	Redis.InitRedis() // 初始化Redis

	ch := make(chan []Commodity2.Commodity, CONSUMER_NUMBER)
	wg.Add(CONSUMER_NUMBER + 1)
	for i := 0; i < CONSUMER_NUMBER; i++ {
		go doRedis(i, ch, &wg)
	}

	go func() {
		sub_wg := sync.WaitGroup{}
		sub_wg.Add(PRODUCT_NUMBER)
		for j := 0; j < PRODUCT_NUMBER; j++ {
			go doMysql(j, ch, &sub_wg)
		}
		sub_wg.Wait()
		close(ch) // 所有的生产者都退出来了， 可以关闭chan了
		wg.Done()
		fmt.Println("all producers exit")
	}()

	wg.Wait()
	fmt.Println("over")
}

func doMysql(j int, ch <-chan []Commodity2.Commodity, wg *sync.WaitGroup) {
	for {
		node, bl := Commodity.Info()
		if !bl {
			break
		} else {
			ch <- node
		}
	}
	wg.Done()
	fmt.Println("producer:", j, "退出了")
}

func doRedis(i int, ch <-chan []Commodity2.Commodity, wg *sync.WaitGroup) {
	for {
		if v, ok := <-ch; ok {
			Commodity3.Consumer(v)
		} else {
			break; //表示channel已经被关闭，退出当前goroutine
		}
	}
	wg.Done()
	fmt.Println("consumer ", i, " 退出了")
}

package main

import (
	"github.com/weiwenwang/ProductAndConsumer/Mysql"
	"github.com/weiwenwang/ProductAndConsumer/Redis"
	Commodity3 "github.com/weiwenwang/ProductAndConsumer/Redis/Commodity"
	Commodity2 "github.com/weiwenwang/ProductAndConsumer/Commodity"
	"github.com/weiwenwang/ProductAndConsumer/Mysql/Commodity"
	"fmt"
	"sync"
	"time"
)

const CONSUMER_NUMBER = 1
const PRODUCT_NUMBER = 20

func main() {
	fmt.Println("begin:", time.Now())
	wg := sync.WaitGroup{}
	Mysql.InitDB()
	Redis.InitRedis()

	ch := make(chan []Commodity2.Commodity, CONSUMER_NUMBER)
	wg.Add(CONSUMER_NUMBER + 1)
	for i := 0; i < CONSUMER_NUMBER; i++ {
		go doRedis(i, ch, &wg)
	}

	go func(wg *sync.WaitGroup) {
		wgg := sync.WaitGroup{}
		wgg.Add(PRODUCT_NUMBER)
		for j := 0; j < PRODUCT_NUMBER; j++ {
			go doMysql(j, ch, &wgg)
		}
		wgg.Wait()
		close(ch)
		wg.Done()
		fmt.Println("product exit")
	}(&wg)

	wg.Wait()
	fmt.Println("over")
	fmt.Println("end:", time.Now())
}

func doMysql(j int, ch chan []Commodity2.Commodity, wg *sync.WaitGroup) {
	for {
		node, bl := Commodity.Info()
		if !bl {
			break
		} else {
			ch <- node
		}
	}
	wg.Done()
	fmt.Println("product:", j)
}

func doRedis(i int, ch chan []Commodity2.Commodity, wg *sync.WaitGroup) {
	for {
		if v, ok := <-ch; ok {
			Commodity3.Consumer(v)
		} else {
			break; //表示channel已经被关闭，退出循环
		}
	}
	wg.Done()
	fmt.Println("goroutine ", i, " 退出了")
}

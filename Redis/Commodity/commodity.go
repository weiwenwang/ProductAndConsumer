package Commodity

import "github.com/weiwenwang/ProductAndConsumer/Commodity"
import (
	"github.com/weiwenwang/ProductAndConsumer/Redis"
	"fmt"
	"strconv"
)

func Consumer(commodities []Commodity.Commodity) {
	Myredis := Redis.RedisClient.Get()
	defer Myredis.Close()
	for _, v := range commodities {
		Myredis.Send("SET", "pkid:"+strconv.Itoa(v.Pkid),
			"Barcode:"+strconv.Itoa(v.Barcode))
	}

	Myredis.Flush()
	_, err := Myredis.Receive()
	if err != nil {
		fmt.Println("redis get failed:", err)
	}

}

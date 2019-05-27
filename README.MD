### 生产者
查询mysql

### 消费者
redis

goroutine直接的通讯用channel

![95B59F05-6E6F-4272-A4D3-8BE2CD48379F.png](https://i.loli.net/2019/05/27/5cebe663dce4e19631.png)

### 其他
- 用到的第三方library, gorm, redigo
- 消费者在拿到多条数据之后， 操作redis使用了管道， 减少io，提高效率
- 消费者redis用了连接池
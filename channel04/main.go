package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生产者消费者模型
//使用goroutine和channel实现一个简易的生产者消费者模型
//生产者：产生随机数
//消费者：计算每个随机数每个位的数字的和

//1个生产者 20个消费者
var itemChan chan *item
var resultChan chan *result

type item struct {
	id  int64
	num int64
}
type result struct {
	item
	sum int64
}

//生产者
func producer(ch chan *item) {
	//1.生成随机数
	var id int64
	for {
		id++
		number := rand.Int63()
		tmp := &item{
			id:  id,
			num: number,
		}
		//2。把随机数发送到通道中
		ch <- tmp
	}

}

//计算一个数字，每一位上的和
func calc(num int64) int64 {
	var sum int64
	for num > 0 {
		sum = sum + num%10
		num = num / 10
	}
	return sum

}

//消费者
func consumer(ch chan *item, resultChan chan *result) {
	for tmp := range ch {
		sum := calc(tmp.num)
		//构造result
		retObj := &result{
			item: *tmp,
			sum:  sum,
		}
		resultChan <- retObj
	}

}

//打印结果
func printResult(resultChan chan *result) {
	for ret := range resultChan {
		fmt.Printf("id:%v,num:%v,sum%v\n", ret.item.id, ret.item.num, ret.sum)
		time.Sleep(time.Second)
	}

}

func startWorker(n int, ch chan *item, resultChan chan *result) {
	for i := 0; i < n; i++ {
		go consumer(ch, resultChan)
	}
}

func main() {
	itemChan = make(chan *item, 100)
	resultChan = make(chan *result, 10)
	go producer(itemChan)
	startWorker(20, itemChan, resultChan)

	//打印结果
	printResult(resultChan)
	// //给rand
	// rand.Seed(time.Now().UnixNano())
	// var ret = rand.Int63() //int64正整数
	// fmt.Println(ret)

}

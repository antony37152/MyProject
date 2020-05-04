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

//itemChan 管道，用于将生产者的随机数传入
var itemChan chan *item

//resultChan管道，用于将结果传入
var resultChan chan *result

//item结构体，存入随机数
type item struct {
	id  int64
	num int64
}

//result结构体 将itemChan中取出的随机数item，
// 与随机各个位数之和sum存入、
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
	//用tmp 遍历取出管道中的随机数
	for tmp := range ch {
		//计算出：取出的随机数的数位和sum
		sum := calc(tmp.num)
		//构造result结构体实例，存入从管道中取出的item,及sum
		retObj := &result{
			item: *tmp,
			sum:  sum,
		}
		//将构造出的result再次传入管道resultChan
		resultChan <- retObj
	}

}

//打印结果
func printResult(resultChan chan *result) {
	//遍历出 管道resultChan 中的元素
	for ret := range resultChan {
		//打印出管道中取出的result中的各个属性
		fmt.Printf("id:%v,num:%v,sum%v\n",
			ret.item.id, ret.item.num, ret.sum)
		time.Sleep(time.Second)
	}

}

//对consumer()进行封装，产生n个consumer()线程
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

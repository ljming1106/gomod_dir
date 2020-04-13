package concurrent

/*
测试并发读数据
1、新建N个goroutine去并发读数据。其中，新建goroutine知add一个waitGroup，goroutine读结束Done（defer），根据一定的规则避免数据竞态
2、读完发回主goroutine中
3、主goroutine新建一个goroutine等待所有的读goroutine结束，然后关闭管道
4、最后merge数据返回
 */

import (
	"fmt"
	"sync"
)

const GETNUM = 10

func GetData() {
	dataChannel := make(chan int)
	var wg sync.WaitGroup
	data := [1000000] int{}
	var res []int
	var num int
	var forNum int
	arrayLen := len(data)
	if arrayLen > GETNUM {
		forNum = GETNUM
	} else {
		forNum = arrayLen
	}

	dealData(forNum, &data, dataChannel ,&wg)

	go func() {//Closer
		wg.Wait()
		close(dataChannel)
	}()

	for v := range dataChannel {
		res = append(res, v)
		num++
	}
	fmt.Println("Num : ", num,len(res))
}


func dealData(forNum int, data *[1000000]int, dataChannel chan int,wg *sync.WaitGroup) {
	for i := 0; i < forNum; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			size := len(data)
			for j := i; j < size; j = j + forNum {
				temp := data[j]
				dataChannel <- temp
			}
		}(i)
	}
}
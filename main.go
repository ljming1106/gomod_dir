package main

import (
"gomod_dir/concurrent"
//"gomod_dir/practice"
)

func main() {
	///*测试并发读数据*/
	concurrent.GetData()
	//
	///*测试并发程序性能*/
	//concurrent.Deal()

	/*测试slice使用*/
	//practice.TestSlice()

	/*channel任务池*/
	//practice.Exec_channel_workpool()
}

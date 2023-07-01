package queue

import (
	"github.com/xgd16/gf-x-rabbitMQ/rabbitMQ"
)

func Service() {
	rabbitMQ.QueueService(map[string]error{
		//"测试队列": rabbitMQ.CreateConsumerHandler(&types.RegisterHandler[types.TestQueueData]{
		//	Handler:    handlers.TestQueue, // 执行的函数
		//	TaskName:   "TestQueue",        // 订阅的任务名称
		//	SyncNum:    30,                 // 同时运行的携程数
		//	FieldNames: []string{"age"},    // 基于哪些字段进行单线程限制
		//}),
	})
}

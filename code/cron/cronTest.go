package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func CronTest() {
	// 创建一个新的 cron 实例
	c := cron.New()

	// 添加一个 cron 作业，使用 cron 表达式来指定任务的执行时间
	c.AddFunc("*/1 * * * *", doSomeThing)

	// 启动 cron 作业调度器
	c.Start()

	time.Sleep(5 * time.Second)

}

func doSomeThing() {
	fmt.Println("do some")
}

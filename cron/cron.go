/*
 * @Description:定时任务执行
 * @Author: Caoxiang
 * @Date: 2021-01-22 14:41:17
 * @LastEditors: Caoxiang
 */
package main

import (
	"time"

	"github.com/newinternetboy/poor_union/pkg/logging"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	// c.AddFunc("* * * * * *", func() {
	// 	logging.Info("Run cron.ReceiveCleanMessageAndCleanCache")
	// })
	// 每隔两秒执行一次
	c.AddFunc("*/2 * * * * *", func() {
		logging.Info("每隔两秒执行一次")
	})
	//每天凌晨十二点执行
	c.AddFunc("0 0 0 * * *", func() {
		logging.Info("每天凌晨十二点执行")
	})
	//每月的十号上午十点发工资
	c.AddFunc("0 0 10 10 * *", func() {
		logging.Info("每月十号上午十点执行")
	})
	c.Start()
	ticker := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			ticker.Reset(time.Second * 10)
		}
	}
}

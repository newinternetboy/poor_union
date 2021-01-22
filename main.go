/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-18 12:01:49
 * @LastEditors: Caoxiang
 */
package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/newinternetboy/poor_union/models"
	"github.com/newinternetboy/poor_union/pkg/logging"
	"github.com/newinternetboy/poor_union/pkg/setting"
	"github.com/newinternetboy/poor_union/routers"
)

func main() {
	//配置初始化
	setting.Setup()
	models.Setup()
	logging.Setup()

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

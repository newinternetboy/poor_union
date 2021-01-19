/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-18 12:01:49
 * @LastEditors: Caoxiang
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/newinternetboy/poor_union/pkg/setting"
	"github.com/newinternetboy/poor_union/routers"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

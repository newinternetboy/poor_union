/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-20 10:55:08
 * @LastEditors: Caoxiang
 */
package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath               = "runtime/logs/"
	LogSaveName               = "log"
	LogFileExt                = "log"
	TimeFormat                = "20060102"
	HourDurationLogTimeFormat = "15"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s/%s/", LogSavePath, time.Now().Format(TimeFormat))
}

func getLogFileFullPath() string {
	return fmt.Sprintf("%s%s.%s", getLogFilePath(), time.Now().Format(HourDurationLogTimeFormat), LogFileExt)
}

func openLogFile(filepath string) *os.File {
	_, err := os.Stat(filepath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}
	handle, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile:%v", err)
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

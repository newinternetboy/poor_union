/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-20 10:55:08
 * @LastEditors: Caoxiang
 */
package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/newinternetboy/poor_union/pkg/file"
	"github.com/newinternetboy/poor_union/pkg/setting"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s/", setting.AppSetting.RuntimeRootPath+setting.AppSetting.LogSavePath, time.Now().Format(setting.AppSetting.TimeFormat))
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s", getLogFilePath(), time.Now().Format(setting.AppSetting.HourDurationLogTimeFormat), setting.AppSetting.LogFileExt)
}

func openLogFile(fileName, filepath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err:%v", err)
	}

	src := dir + "/" + filepath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	f, err := file.Open(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

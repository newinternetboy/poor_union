/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-19 14:35:33
 * @LastEditors: Caoxiang
 */
package util

import (
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/newinternetboy/poor_union/pkg/logging"
)

func ValidErrorPrintf(errors []*validation.Error) {
	for _, err := range errors {
		logging.Error(fmt.Sprintf("err.Key:%s;err.Message:%s", err.Key, err.Message))
	}
}

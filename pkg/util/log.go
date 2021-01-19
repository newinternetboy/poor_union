/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-19 14:35:33
 * @LastEditors: Caoxiang
 */
package util

import (
	"log"

	"github.com/astaxie/beego/validation"
)

func ValidErrorPrintf(errors []*validation.Error) {
	for _, err := range errors {
		log.Printf("err.Key:%s;err.Message:%s", err.Key, err.Message)
	}
}

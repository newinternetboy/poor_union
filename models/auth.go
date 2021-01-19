/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-01-19 16:31:49
 * @LastEditors: Caoxiang
 */
package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//校验用户和密码是否正确
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	} else {
		return false
	}
}

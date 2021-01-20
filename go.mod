module github.com/newinternetboy/poor_union

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/ugorji/go v1.2.3 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sys v0.0.0-20210113181707-4bcb84eeeb78 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// //未推送到远端的包依赖指向本地
// replace (
// 	github.com/newinternetboy/poor_union/conf => /Users/caoxiang/caoxiang/go/poor_union/conf
// 	github.com/newinternetboy/poor_union/models => /Users/caoxiang/caoxiang/go/poor_union/models
// 	github.com/newinternetboy/poor_union/pkg/e => /Users/caoxiang/caoxiang/go/poor_union/pkg/e
// 	github.com/newinternetboy/poor_union/pkg/setting => /Users/caoxiang/caoxiang/go/poor_union/pkg/setting
// 	github.com/newinternetboy/poor_union/routers => /Users/caoxiang/caoxiang/go/poor_union/routers
// 	github.com/newinternetboy/poor_union/setting => /Users/caoxiang/caoxiang/go/poor_union/setting
// 	github.com/newinternetboy/poor_union/pkg/logging => /Users/caoxiang/caoxiang/go/poor_union/pkg/logging
// )

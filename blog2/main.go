package main

import (
	"GOproject/blog2/common"
	"GOproject/blog2/routes"
	_ "github.com/go-sql-driver/mysql"

	"net/http"
)

func main() {
	// 获取初始化的数据库
	db := common.InitDB()
	// 延迟关闭数据库
	defer db.Close()
	// 创建路由引擎
	r := routes.NewRouter()
	// 配置静态文件路径
	r.StaticFS("/image", http.Dir("./static/image"))
	// 启动路由
	panic(r.Run(":8080"))
}

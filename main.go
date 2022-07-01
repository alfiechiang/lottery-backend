package main

import (
	"lottery/conf"
	"lottery/msg"
	"lottery/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()
	msg.Init()
	// 装载路由
	r := server.NewRouter()
	r.Run(":8080")
}

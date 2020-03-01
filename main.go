package main

import (
	"github.com/gogf/gf/frame/g"
	_ "yj-app/boot"
	_ "yj-app/router"
)

// @title 云捷GO 自动生成API文档
// @version 1.0
// @description <a href="/tool/swagger?a=r">重新生成文档</a>

// @host localhost
// @BasePath /api
func main() {
	g.Server().Start()

	api := g.Server("api")
	api.SetPort(8080)
	api.Start()
	g.Wait()
}

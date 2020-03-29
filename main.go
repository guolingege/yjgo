package main

import (
	"github.com/gogf/gf/frame/g"
	_ "yj-app/boot"
	_ "yj-app/router"
)

// @title 云捷GO 自动生成API文档
// @version 1.0
// @description 生成文档请在调试模式下进行<a href="/tool/swagger?a=r">重新生成文档</a>

// @host localhost
// @BasePath /api
func main() {

	serverSwitch := g.Cfg().GetBool("status.admin")
	apiSwitch := g.Cfg().GetBool("status.api")

	if serverSwitch {
		g.Server().Start()
	}

	if apiSwitch {
		address := g.Cfg().GetString("api.Address")
		api := g.Server("api")
		api.SetAddr(address)
		api.Start()
	}

	g.Wait()
}

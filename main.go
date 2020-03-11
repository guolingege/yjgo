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

	serverSwitch := g.Cfg().GetBool("switch.server")
	apiSwitch := g.Cfg().GetBool("switch.api")

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

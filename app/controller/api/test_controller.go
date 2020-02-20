package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
	"yj-app/app/service/utils/response"
)

// @Summary api测试
// @Description api测试
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CommonRes
// @Router /api/index [get]
func Index(r *ghttp.Request) {
	response.SucessData(r, "api测试", g.Map{"test": "test"}, model.Buniss_Other, "Index")
}

package boot

import (
	"github.com/gogf/gf/os/gview"
	"yj-app/app/service/system/config"
	"yj-app/app/service/system/dict"
	"yj-app/app/service/system/permission"
)

func init() {
	gview.Instance().BindFuncMap(gview.FuncMap{
		"hasPermi":          permission.HasPermi,
		"getPermiButton":    permission.GetPermiButton,
		"getDictLabel":      dict.GetDictLabel,
		"getDictTypeSelect": dict.GetDictTypeSelect,
		"getDictTypeRadio":  dict.GetDictTypeRadio,
		"getDictTypeData":   dict.GetDictTypeData,
		"ossUrl":            config.GetOssUrl,
	})
}

package ip

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}

	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}

	url := "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
	bytes := ghttp.GetBytes(url)
	json, err := gjson.DecodeToJson(bytes)
	if err != nil {
		return ""
	}
	if json.GetInt("code") == 0 {
		return json.GetString("data.city")
	} else {
		return ""
	}
}

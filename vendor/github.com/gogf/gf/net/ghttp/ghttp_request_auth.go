// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gogf/gf/encoding/gbase64"
)

// 设置Basic Auth校验提示
func (r *Request) setBasicAuth(tips ...string) {
	realm := ""
	if len(tips) > 0 && tips[0] != "" {
		realm = tips[0]
	} else {
		realm = "Need Login"
	}
	r.Response.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	r.Response.WriteHeader(http.StatusUnauthorized)
}

// 设置HTTP基础账号密码认证，如果用户没有提交账号密码，那么提示用户输出信息。
// 验证成功之后返回true，否则返回false。
func (r *Request) BasicAuth(user, pass string, tips ...string) bool {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		r.setBasicAuth(tips...)
		return false
	}
	authArray := strings.SplitN(auth, " ", 2)
	if len(authArray) != 2 {
		r.Response.WriteStatus(http.StatusForbidden)
		return false
	}
	switch authArray[0] {
	case "Basic":
		authBytes, err := gbase64.DecodeString(authArray[1])
		if err != nil {
			r.Response.WriteStatus(http.StatusForbidden, err.Error())
			return false
		}
		authArray := strings.SplitN(string(authBytes), ":", 2)
		if len(authArray) != 2 {
			r.Response.WriteStatus(http.StatusForbidden)
			return false
		}
		if authArray[0] != user || authArray[1] != pass {
			r.setBasicAuth(tips...)
			return false
		}
		return true

	default:
		r.Response.WriteStatus(http.StatusForbidden)
		return false
	}
	return false
}

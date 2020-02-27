// Copyright 2018 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"fmt"
	"github.com/gogf/gf/errors/gerror"
)

const (
	gPATH_FILTER_KEY = "github.com/gogf/gf/"
)

// 处理服务错误信息，主要是panic，http请求的status由access log进行管理
func (s *Server) handleAccessLog(r *Request) {
	if !s.IsAccessLogEnabled() {
		return
	}
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	s.config.Logger.File(s.config.AccessLogPattern).
		Stdout(s.config.LogStdout).
		Printf(
			`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s"`,
			r.Response.Status,
			r.Method, scheme, r.Host, r.URL.String(), r.Proto,
			float64(r.LeaveTime-r.EnterTime)/1000,
			r.GetClientIp(), r.Referer(), r.UserAgent(),
		)
}

// 处理服务错误信息，主要是panic，http请求的status由access log进行管理
func (s *Server) handleErrorLog(err error, r *Request) {
	// 错误输出默认是开启的
	if !s.IsErrorLogEnabled() {
		return
	}

	// 错误日志信息
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	content := fmt.Sprintf(`%d, "%s %s %s %s %s" %.3f, %s, "%s", "%s"`,
		r.Response.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime-r.EnterTime)/1000,
		r.GetClientIp(), r.Referer(), r.UserAgent(),
	)
	if stack := gerror.Stack(err); stack != "" {
		content += "\nStack:\n" + stack
		s.config.Logger.File(s.config.AccessLogPattern).
			Stack(false).
			Stdout(s.config.LogStdout).
			Error(content)
		return
	}
	s.config.Logger.File(s.config.AccessLogPattern).
		Stack(s.config.ErrorStack).
		StackWithFilter(gPATH_FILTER_KEY).
		Stdout(s.config.LogStdout).
		Error(content)
}

// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"fmt"
	"github.com/gogf/gf/os/gres"
	"github.com/gogf/gf/os/gview"
	"net/http"
	"strings"

	"github.com/gogf/gf/os/gsession"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gregex"
)

// Request is the context object for a request.
type Request struct {
	*http.Request
	Server          *Server                // Parent server.
	Cookie          *Cookie                // Cookie.
	Session         *gsession.Session      // Session.
	Response        *Response              // Corresponding Response of this request.
	Router          *Router                // Matched Router for this request. Note that it's only available in HTTP handler, not in HOOK or MiddleWare.
	EnterTime       int64                  // Request starting time in microseconds.
	LeaveTime       int64                  // Request ending time in microseconds.
	Middleware      *Middleware            // The middleware manager.
	StaticFile      *StaticFile            // Static file object when static file serving.
	handlers        []*handlerParsedItem   // All matched handlers containing handler, hook and middleware for this request .
	hasHookHandler  bool                   // A bool marking whether there's hook handler in the handlers for performance purpose.
	hasServeHandler bool                   // A bool marking whether there's serving handler in the handlers for performance purpose.
	parsedQuery     bool                   // A bool marking whether the GET parameters parsed.
	parsedBody      bool                   // A bool marking whether the request body parsed.
	parsedForm      bool                   // A bool marking whether request Form parsed for HTTP method PUT, POST, PATCH.
	paramsMap       map[string]interface{} // Custom parameters.
	routerMap       map[string]string      // Router parameters map, which might be nil if there're no router parameters.
	queryMap        map[string]interface{} // Query parameters map, which is nil if there's no query string.
	formMap         map[string]interface{} // Form parameters map, which is nil if there's no form data from client.
	bodyMap         map[string]interface{} // Body parameters map, which might be nil if there're no body content.
	error           error                  // Current executing error of the request.
	exit            bool                   // A bool marking whether current request is exited.
	parsedHost      string                 // The parsed host name for current host used by GetHost function.
	clientIp        string                 // The parsed client ip for current host used by GetClientIp function.
	bodyContent     []byte                 // Request body content.
	isFileRequest   bool                   // A bool marking whether current request is file serving.
	viewObject      *gview.View            // Custom template view engine object for this response.
	viewParams      gview.Params           // Custom template view variables for this response.
}

// StaticFile is the file struct for static file service.
type StaticFile struct {
	File  *gres.File // Resource file object.
	Path  string     // File path.
	IsDir bool       // Is directory.
}

// newRequest creates and returns a new request object.
func newRequest(s *Server, r *http.Request, w http.ResponseWriter) *Request {
	request := &Request{
		Server:    s,
		Request:   r,
		Response:  newResponse(s, w),
		EnterTime: gtime.TimestampMicro(),
	}
	request.Cookie = GetCookie(request)
	request.Session = s.sessionManager.New(request.GetSessionId())
	request.Response.Request = request
	request.Middleware = &Middleware{
		request: request,
	}
	return request
}

// WebSocket upgrades current request as a websocket request.
// It returns a new WebSocket object if success, or the error if failure.
// Note that the request should be a websocket request, or it will surely fail upgrading.
func (r *Request) WebSocket() (*WebSocket, error) {
	if conn, err := wsUpgrader.Upgrade(r.Response.Writer, r.Request, nil); err == nil {
		return &WebSocket{
			conn,
		}, nil
	} else {
		return nil, err
	}
}

// Exit exits executing of current HTTP handler.
func (r *Request) Exit() {
	panic(gEXCEPTION_EXIT)
}

// ExitAll exits executing of current and following HTTP handlers.
func (r *Request) ExitAll() {
	r.exit = true
	panic(gEXCEPTION_EXIT_ALL)
}

// ExitHook exits executing of current and following HTTP HOOK handlers.
func (r *Request) ExitHook() {
	panic(gEXCEPTION_EXIT_HOOK)
}

// IsExited checks and returns whether current request is exited.
func (r *Request) IsExited() bool {
	return r.exit
}

// GetHost returns current request host name, which might be a domain or an IP without port.
func (r *Request) GetHost() string {
	if len(r.parsedHost) == 0 {
		array, _ := gregex.MatchString(`(.+):(\d+)`, r.Host)
		if len(array) > 1 {
			r.parsedHost = array[1]
		} else {
			r.parsedHost = r.Host
		}
	}
	return r.parsedHost
}

// IsFileRequest checks and returns whether current request is serving file.
func (r *Request) IsFileRequest() bool {
	return r.isFileRequest
}

// IsAjaxRequest checks and returns whether current request is an AJAX request.
func (r *Request) IsAjaxRequest() bool {
	return strings.EqualFold(r.Header.Get("X-Requested-With"), "XMLHttpRequest")
}

// GetClientIp returns the client ip of this request without port.
func (r *Request) GetClientIp() string {
	if len(r.clientIp) == 0 {
		if r.clientIp = r.Header.Get("X-Real-IP"); r.clientIp == "" {
			array, _ := gregex.MatchString(`(.+):(\d+)`, r.RemoteAddr)
			if len(array) > 1 {
				r.clientIp = array[1]
			} else {
				r.clientIp = r.RemoteAddr
			}
		}
	}
	return r.clientIp
}

// GetUrl returns current URL of this request.
func (r *Request) GetUrl() string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf(`%s://%s%s`, scheme, r.Host, r.URL.String())
}

// GetSessionId retrieves and returns session id from cookie or header.
func (r *Request) GetSessionId() string {
	id := r.Cookie.GetSessionId()
	if id == "" {
		id = r.Header.Get(r.Server.GetSessionIdName())
	}
	return id
}

// GetReferer returns referer of this request.
func (r *Request) GetReferer() string {
	return r.Header.Get("Referer")
}

// GetError returns the error occurs in the procedure of the request.
// It returns nil if there's no error.
func (r *Request) GetError() error {
	return r.error
}

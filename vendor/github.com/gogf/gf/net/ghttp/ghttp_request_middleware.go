// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package ghttp

import (
	"net/http"
	"reflect"

	"github.com/gogf/gf/errors/gerror"

	"github.com/gogf/gf/util/gutil"
)

// Middleware is the plugin for request workflow management.
type Middleware struct {
	served         bool     // Is the request served, which is used for checking response status 404.
	request        *Request // The request object pointer.
	handlerIndex   int      // Index number for executing sequence purpose for handler items.
	handlerMDIndex int      // Index number for executing sequence purpose for bound middleware of handler item.
}

// Next calls the next workflow handler.
func (m *Middleware) Next() {
	var item *handlerParsedItem
	var loop = true
	for loop {
		// Check whether the request is exited.
		if m.request.IsExited() || m.handlerIndex >= len(m.request.handlers) {
			break
		}
		item = m.request.handlers[m.handlerIndex]
		// Filter the HOOK handlers, which are designed to be called in another standalone procedure.
		if item.handler.itemType == gHANDLER_TYPE_HOOK {
			m.handlerIndex++
			continue
		}
		// Router values switching.
		m.request.routerMap = item.values
		// Current router switching.
		m.request.Router = item.handler.router

		gutil.TryCatch(func() {
			// Execute bound middleware array of the item if it's not empty.
			if m.handlerMDIndex < len(item.handler.middleware) {
				md := item.handler.middleware[m.handlerMDIndex]
				m.handlerMDIndex++
				niceCallFunc(func() {
					md(m.request)
				})
				loop = false
				return
			}
			m.handlerIndex++

			switch item.handler.itemType {
			// Service controller.
			case gHANDLER_TYPE_CONTROLLER:
				m.served = true
				if m.request.IsExited() {
					break
				}
				c := reflect.New(item.handler.ctrlInfo.reflect)
				niceCallFunc(func() {
					c.MethodByName("Init").Call([]reflect.Value{reflect.ValueOf(m.request)})
				})
				if !m.request.IsExited() {
					niceCallFunc(func() {
						c.MethodByName(item.handler.ctrlInfo.name).Call(nil)
					})
				}
				if !m.request.IsExited() {
					niceCallFunc(func() {
						c.MethodByName("Shut").Call(nil)
					})
				}

			// Service object.
			case gHANDLER_TYPE_OBJECT:
				m.served = true
				if m.request.IsExited() {
					break
				}
				if item.handler.initFunc != nil {
					niceCallFunc(func() {
						item.handler.initFunc(m.request)
					})
				}
				if !m.request.IsExited() {
					niceCallFunc(func() {
						item.handler.itemFunc(m.request)
					})
				}
				if !m.request.IsExited() && item.handler.shutFunc != nil {
					niceCallFunc(func() {
						item.handler.shutFunc(m.request)
					})
				}

			// Service handler.
			case gHANDLER_TYPE_HANDLER:
				m.served = true
				if m.request.IsExited() {
					break
				}
				niceCallFunc(func() {
					item.handler.itemFunc(m.request)
				})

			// Global middleware array.
			case gHANDLER_TYPE_MIDDLEWARE:
				niceCallFunc(func() {
					item.handler.itemFunc(m.request)
				})
				// It does not continue calling next middleware after another middleware done.
				// There should be a "Next" function to be called in the middleware in order to manage the workflow.
				loop = false
			}
		}, func(exception interface{}) {
			m.request.error = gerror.Newf("%v", exception)
			m.request.Response.WriteStatus(http.StatusInternalServerError, exception)
		})
	}
	// Check the http status code after all handler and middleware done.
	if m.request.IsExited() || m.handlerIndex >= len(m.request.handlers) {
		if m.request.Response.Status == 0 {
			if m.request.Middleware.served {
				m.request.Response.WriteHeader(http.StatusOK)
			} else {
				m.request.Response.WriteHeader(http.StatusNotFound)
			}
		}
	}
}

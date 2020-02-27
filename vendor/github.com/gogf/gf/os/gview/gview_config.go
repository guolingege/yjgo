// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gview

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/i18n/gi18n"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gres"
	"github.com/gogf/gf/os/gspath"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
)

// Config is the configuration object for template engine.
type Config struct {
	Paths       []string               // Searching array for path, NOT concurrent-safe for performance purpose.
	Data        map[string]interface{} // Global template variables.
	DefaultFile string                 // Default template file for parsing.
	Delimiters  []string               // Custom template delimiters.
	AutoEncode  bool                   // Automatically encodes and provides safe html output, which is good for avoiding XSS.
}

// SetConfig sets the configuration for view.
func (view *View) SetConfig(config Config) error {
	var err error
	if len(config.Paths) > 0 {
		for _, v := range config.Paths {
			if err = view.AddPath(v); err != nil {
				return err
			}
		}
	}
	if len(config.Data) > 0 {
		view.Assigns(config.Data)
	}
	if config.DefaultFile != "" {
		view.SetDefaultFile(config.DefaultFile)
	}
	if len(config.Delimiters) > 1 {
		view.SetDelimiters(config.Delimiters[0], config.Delimiters[1])
	}
	view.config = config
	// Clear global template object cache.
	// It's just cache, do not hesitate clearing it.
	templates.Clear()
	return nil
}

// SetConfigWithMap set configurations with map for the view.
func (view *View) SetConfigWithMap(m map[string]interface{}) error {
	if m == nil || len(m) == 0 {
		return errors.New("configuration cannot be empty")
	}
	// Most common used configuration support for single view path.
	_, v1 := gutil.MapPossibleItemByKey(m, "paths")
	_, v2 := gutil.MapPossibleItemByKey(m, "path")
	if v1 == nil && v2 != nil {
		m["paths"] = []interface{}{v2}
	}
	config := Config{}
	err := gconv.Struct(m, &config)
	if err != nil {
		return err
	}
	return view.SetConfig(config)
}

// SetPath sets the template directory path for template file search.
// The parameter <path> can be absolute or relative path, but absolute path is suggested.
func (view *View) SetPath(path string) error {
	isDir := false
	realPath := ""
	if file := gres.Get(path); file != nil {
		realPath = path
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		realPath = gfile.RealPath(path)
		if realPath == "" {
			// Relative path.
			view.paths.RLockFunc(func(array []string) {
				for _, v := range array {
					if path, _ := gspath.Search(v, path); path != "" {
						realPath = path
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = gfile.IsDir(realPath)
		}
	}
	// Path not exist.
	if realPath == "" {
		err := errors.New(fmt.Sprintf(`[gview] SetPath failed: path "%s" does not exist`, path))
		if errorPrint() {
			glog.Error(err)
		}
		return err
	}
	// Should be a directory.
	if !isDir {
		err := errors.New(fmt.Sprintf(`[gview] SetPath failed: path "%s" should be directory type`, path))
		if errorPrint() {
			glog.Error(err)
		}
		return err
	}
	// Repeated path adding check.
	if view.paths.Search(realPath) != -1 {
		return nil
	}
	view.paths.Clear()
	view.paths.Append(realPath)
	view.fileCacheMap.Clear()
	//glog.Debug("[gview] SetPath:", realPath)
	return nil
}

// AddPath adds a absolute or relative path to the search paths.
func (view *View) AddPath(path string) error {
	isDir := false
	realPath := ""
	if file := gres.Get(path); file != nil {
		realPath = path
		isDir = file.FileInfo().IsDir()
	} else {
		// Absolute path.
		realPath = gfile.RealPath(path)
		if realPath == "" {
			// Relative path.
			view.paths.RLockFunc(func(array []string) {
				for _, v := range array {
					if path, _ := gspath.Search(v, path); path != "" {
						realPath = path
						break
					}
				}
			})
		}
		if realPath != "" {
			isDir = gfile.IsDir(realPath)
		}
	}
	// Path not exist.
	if realPath == "" {
		err := errors.New(fmt.Sprintf(`[gview] AddPath failed: path "%s" does not exist`, path))
		if errorPrint() {
			glog.Error(err)
		}
		return err
	}
	// realPath should be type of folder.
	if !isDir {
		err := errors.New(fmt.Sprintf(`[gview] AddPath failed: path "%s" should be directory type`, path))
		if errorPrint() {
			glog.Error(err)
		}
		return err
	}
	// Repeated path adding check.
	if view.paths.Search(realPath) != -1 {
		return nil
	}
	view.paths.Append(realPath)
	view.fileCacheMap.Clear()
	return nil
}

// Assigns binds multiple global template variables to current view object.
// Note that it's not concurrent-safe, which means it would panic
// if it's called in multiple goroutines in runtime.
func (view *View) Assigns(data Params) {
	for k, v := range data {
		view.data[k] = v
	}
}

// Assign binds a global template variable to current view object.
// Note that it's not concurrent-safe, which means it would panic
// if it's called in multiple goroutines in runtime.
func (view *View) Assign(key string, value interface{}) {
	view.data[key] = value
}

// SetDefaultFile sets default template file for parsing.
func (view *View) SetDefaultFile(file string) {
	view.defaultFile = file
}

// SetDelimiters sets customized delimiters for template parsing.
func (view *View) SetDelimiters(left, right string) {
	view.delimiters[0] = left
	view.delimiters[1] = right
}

// SetAutoEncode enables/disables automatically html encoding feature.
// When AutoEncode feature is enables, view engine automatically encodes and provides safe html output,
// which is good for avoid XSS.
func (view *View) SetAutoEncode(enable bool) {
	view.config.AutoEncode = enable
}

// BindFunc registers customized global template function named <name>
// with given function <function> to current view object.
// The <name> is the function name which can be called in template content.
func (view *View) BindFunc(name string, function interface{}) {
	view.funcMap[name] = function
	// Clear global template object cache.
	templates.Clear()
}

// BindFuncMap registers customized global template functions by map to current view object.
// The key of map is the template function name
// and the value of map is the address of customized function.
func (view *View) BindFuncMap(funcMap FuncMap) {
	for k, v := range funcMap {
		view.funcMap[k] = v
	}
	// Clear global template object cache.
	templates.Clear()
}

// SetI18n binds i18n manager to current view engine.
func (view *View) SetI18n(manager *gi18n.Manager) {
	view.i18nManager = manager
}

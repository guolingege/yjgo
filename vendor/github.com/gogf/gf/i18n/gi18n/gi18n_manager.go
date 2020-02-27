// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gi18n

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/gogf/gf/os/glog"

	"github.com/gogf/gf/os/gfsnotify"

	"github.com/gogf/gf/text/gregex"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/encoding/gjson"

	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gres"
)

// Manager, it is concurrent safe, supporting hot reload.
type Manager struct {
	mu      sync.RWMutex
	data    map[string]map[string]string // Translating map.
	pattern string                       // Pattern for regex parsing.
	options Options                      // configuration options.
}

type Options struct {
	Path       string   // I18n files storage path.
	Language   string   // Local language.
	Delimiters []string // Delimiters for variable parsing.
}

var (
	defaultDelimiters = []string{"{#", "}"}
)

// New creates and returns a new i18n manager.
func New(options ...Options) *Manager {
	var opts Options
	if len(options) > 0 {
		opts = options[0]
	} else {
		opts = DefaultOptions()
	}
	if len(opts.Delimiters) == 0 {
		opts.Delimiters = defaultDelimiters
	}
	return &Manager{
		options: opts,
		pattern: fmt.Sprintf(
			`%s(\w+)%s`,
			gregex.Quote(opts.Delimiters[0]),
			gregex.Quote(opts.Delimiters[1]),
		),
	}
}

// DefaultOptions returns the default options for i18n manager.
func DefaultOptions() Options {
	path := "i18n"
	realPath, _ := gfile.Search(path)
	if realPath != "" {
		path = realPath
		// To avoid of the source of GF: github.com/gogf/i18n/gi18n
		if gfile.Exists(path + gfile.Separator + "gi18n") {
			path = ""
		}
	}
	return Options{
		Path:       path,
		Delimiters: []string{"{#", "}"},
	}
}

// SetPath sets the directory path storing i18n files.
func (m *Manager) SetPath(path string) error {
	if gres.Contains(path) {
		m.options.Path = path
	} else {
		realPath, _ := gfile.Search(path)
		if realPath == "" {
			return errors.New(fmt.Sprintf(`%s does not exist`, path))
		}
		m.options.Path = realPath
	}
	return nil
}

// SetLanguage sets the language for translator.
func (m *Manager) SetLanguage(language string) {
	m.options.Language = language
}

// SetDelimiters sets the delimiters for translator.
func (m *Manager) SetDelimiters(left, right string) {
	m.pattern = fmt.Sprintf(`%s(\w+)%s`, gregex.Quote(left), gregex.Quote(right))
}

// T is alias of Translate.
func (m *Manager) T(content string, language ...string) string {
	return m.Translate(content, language...)
}

// Translate translates <content> with configured language.
// The parameter <language> specifies custom translation language ignoring configured language.
func (m *Manager) Translate(content string, language ...string) string {
	m.init()
	m.mu.RLock()
	defer m.mu.RUnlock()
	var data map[string]string
	if len(language) > 0 && language[0] != "" {
		data = m.data[language[0]]
	} else {
		data = m.data[m.options.Language]
	}
	if data == nil {
		return content
	}
	// Parse content as name.
	if v, ok := data[content]; ok {
		return v
	}
	// Parse content as variables container.
	result, _ := gregex.ReplaceStringFuncMatch(m.pattern, content, func(match []string) string {
		if v, ok := data[match[1]]; ok {
			return v
		}
		return match[0]
	})
	return result
}

func (m *Manager) init() {
	m.mu.RLock()
	if m.data != nil {
		m.mu.RUnlock()
		return
	}
	m.mu.RUnlock()

	m.mu.Lock()
	defer m.mu.Unlock()
	if gres.Contains(m.options.Path) {
		files := gres.ScanDirFile(m.options.Path, "*.*", true)
		if len(files) > 0 {
			var path string
			var name string
			var lang string
			var array []string
			m.data = make(map[string]map[string]string)
			for _, file := range files {
				name = file.Name()
				path = name[len(m.options.Path)+1:]
				array = strings.Split(path, "/")
				if len(array) > 1 {
					lang = array[0]
				} else {
					lang = gfile.Name(array[0])
				}
				if m.data[lang] == nil {
					m.data[lang] = make(map[string]string)
				}
				if j, err := gjson.LoadContent(file.Content()); err == nil {
					for k, v := range j.ToMap() {
						m.data[lang][k] = gconv.String(v)
					}
				} else {
					glog.Errorf("load i18n file '%s' failed: %v", name, err)
				}
			}
		}
	} else if m.options.Path != "" {
		files, _ := gfile.ScanDirFile(m.options.Path, "*.*", true)
		if len(files) > 0 {
			var path string
			var lang string
			var array []string
			m.data = make(map[string]map[string]string)
			for _, file := range files {
				path = file[len(m.options.Path)+1:]
				array = strings.Split(path, gfile.Separator)
				if len(array) > 1 {
					lang = array[0]
				} else {
					lang = gfile.Name(array[0])
				}
				if m.data[lang] == nil {
					m.data[lang] = make(map[string]string)
				}
				if j, err := gjson.LoadContent(gfile.GetBytes(file)); err == nil {
					for k, v := range j.ToMap() {
						m.data[lang][k] = gconv.String(v)
					}
				} else {
					glog.Errorf("load i18n file '%s' failed: %v", file, err)
				}
			}
			_, _ = gfsnotify.Add(path, func(event *gfsnotify.Event) {
				m.mu.Lock()
				m.data = nil
				m.mu.Unlock()
				gfsnotify.Exit()
			})
		}
	}
}

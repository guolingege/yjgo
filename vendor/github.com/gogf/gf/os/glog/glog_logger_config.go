// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package glog

import (
	"errors"
	"github.com/gogf/gf/internal/intlog"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
	"io"
	"strings"
)

// Config is the configuration object for logger.
type Config struct {
	Writer      io.Writer // Customized io.Writer.
	Flags       int       // Extra flags for logging output features.
	Path        string    // Logging directory path.
	File        string    // Format for logging file.
	Level       int       // Output level.
	Prefix      string    // Prefix string for every logging content.
	StSkip      int       // Skip count for stack.
	StStatus    int       // Stack status(1: enabled - default; 0: disabled)
	StFilter    string    // Stack string filter.
	HeaderPrint bool      `c:"header"` // Print header or not(true in default).
	StdoutPrint bool      `c:"stdout"` // Output to stdout or not(true in default).
}

// DefaultConfig returns the default configuration for logger.
func DefaultConfig() Config {
	c := Config{
		File:        gDEFAULT_FILE_FORMAT,
		Flags:       F_TIME_STD,
		Level:       LEVEL_ALL,
		StStatus:    1,
		HeaderPrint: true,
		StdoutPrint: true,
	}
	if !defaultDebug {
		c.Level = c.Level & ^LEVEL_DEBU
	}
	return c
}

// SetConfig set configurations for the logger.
func (l *Logger) SetConfig(config Config) error {
	l.config = config
	// Necessary validation.
	if config.Path != "" {
		if err := l.SetPath(config.Path); err != nil {
			intlog.Error(err)
			return err
		}
	}
	intlog.Print(l.config)
	return nil
}

// SetConfigWithMap set configurations with map for the logger.
func (l *Logger) SetConfigWithMap(m map[string]interface{}) error {
	if m == nil || len(m) == 0 {
		return errors.New("configuration cannot be empty")
	}
	// Change string configuration to int value for level.
	levelKey, levelValue := gutil.MapPossibleItemByKey(m, "level")
	if levelValue != nil {
		switch gconv.String(levelValue) {
		case "all":
			m[levelKey] = LEVEL_ALL
		case "dev":
			m[levelKey] = LEVEL_DEV
		case "prod":
			m[levelKey] = LEVEL_PROD
		}
	}
	config := DefaultConfig()
	err := gconv.Struct(m, &config)
	if err != nil {
		return err
	}
	return l.SetConfig(config)
}

// SetLevel sets the logging level.
func (l *Logger) SetLevel(level int) {
	l.config.Level = level
}

// GetLevel returns the logging level value.
func (l *Logger) GetLevel() int {
	return l.config.Level
}

// SetDebug enables/disables the debug level for logger.
// The debug level is enabled in default.
func (l *Logger) SetDebug(debug bool) {
	if debug {
		l.config.Level = l.config.Level | LEVEL_DEBU
	} else {
		l.config.Level = l.config.Level & ^LEVEL_DEBU
	}
}

// SetAsync enables/disables async logging output feature.
func (l *Logger) SetAsync(enabled bool) {
	if enabled {
		l.config.Flags = l.config.Flags | F_ASYNC
	} else {
		l.config.Flags = l.config.Flags & ^F_ASYNC
	}
}

// SetFlags sets extra flags for logging output features.
func (l *Logger) SetFlags(flags int) {
	l.config.Flags = flags
}

// GetFlags returns the flags of logger.
func (l *Logger) GetFlags() int {
	return l.config.Flags
}

// SetStack enables/disables the stack feature in failure logging outputs.
func (l *Logger) SetStack(enabled bool) {
	if enabled {
		l.config.StStatus = 1
	} else {
		l.config.StStatus = 0
	}
}

// SetStackSkip sets the stack offset from the end point.
func (l *Logger) SetStackSkip(skip int) {
	l.config.StSkip = skip
}

// SetStackFilter sets the stack filter from the end point.
func (l *Logger) SetStackFilter(filter string) {
	l.config.StFilter = filter
}

// SetWriter sets the customized logging <writer> for logging.
// The <writer> object should implements the io.Writer interface.
// Developer can use customized logging <writer> to redirect logging output to another service,
// eg: kafka, mysql, mongodb, etc.
func (l *Logger) SetWriter(writer io.Writer) {
	l.config.Writer = writer
}

// GetWriter returns the customized writer object, which implements the io.Writer interface.
// It returns nil if no writer previously set.
func (l *Logger) GetWriter() io.Writer {
	return l.config.Writer
}

// SetPath sets the directory path for file logging.
func (l *Logger) SetPath(path string) error {
	if path == "" {
		return errors.New("logging path is empty")
	}
	if !gfile.Exists(path) {
		if err := gfile.Mkdir(path); err != nil {
			//fmt.Fprintln(os.Stderr, fmt.Sprintf(`[glog] mkdir "%s" failed: %s`, path, err.Error()))
			return err
		}
	}
	l.config.Path = strings.TrimRight(path, gfile.Separator)
	return nil
}

// GetPath returns the logging directory path for file logging.
// It returns empty string if no directory path set.
func (l *Logger) GetPath() string {
	return l.config.Path
}

// SetFile sets the file name <pattern> for file logging.
// Datetime pattern can be used in <pattern>, eg: access-{Ymd}.log.
// The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log
func (l *Logger) SetFile(pattern string) {
	l.config.File = pattern
}

// SetStdoutPrint sets whether output the logging contents to stdout, which is true in default.
func (l *Logger) SetStdoutPrint(enabled bool) {
	l.config.StdoutPrint = enabled
}

// SetHeaderPrint sets whether output header of the logging contents, which is true in default.
func (l *Logger) SetHeaderPrint(enabled bool) {
	l.config.HeaderPrint = enabled
}

// SetPrefix sets prefix string for every logging content.
// Prefix is part of header, which means if header output is shut, no prefix will be output.
func (l *Logger) SetPrefix(prefix string) {
	l.config.Prefix = prefix
}

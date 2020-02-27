// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gres

import (
	"archive/zip"
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/encoding/gcompress"
	"github.com/gogf/gf/os/gfile"
)

const (
	gPACKAGE_TEMPLATE = `package %s

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("%s"); err != nil {
		panic(err)
	}
}
`
)

// Pack packs the path specified by <srcPaths> into bytes.
// The unnecessary parameter <keyPrefix> indicates the prefix for each file
// packed into the result bytes.
//
// Note that parameter <srcPaths> supports multiple paths join with ','.
func Pack(srcPaths string, keyPrefix ...string) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	headerPrefix := ""
	if len(keyPrefix) > 0 && keyPrefix[0] != "" {
		headerPrefix = keyPrefix[0]
	}
	err := gcompress.ZipPathWriter(srcPaths, buffer, headerPrefix)
	if err != nil {
		return nil, err
	}
	// Gzip the data bytes to reduce the size.
	return gcompress.Gzip(buffer.Bytes(), 9)
}

// PackToFile packs the path specified by <srcPaths> to target file <dstPath>.
// The unnecessary parameter <keyPrefix> indicates the prefix for each file
// packed into the result bytes.
//
// Note that parameter <srcPaths> supports multiple paths join with ','.
func PackToFile(srcPaths, dstPath string, keyPrefix ...string) error {
	data, err := Pack(srcPaths, keyPrefix...)
	if err != nil {
		return err
	}
	return gfile.PutBytes(dstPath, data)
}

// PackToGoFile packs the path specified by <srcPaths> to target go file <goFilePath>
// with given package name <pkgName>.
//
// The unnecessary parameter <keyPrefix> indicates the prefix for each file
// packed into the result bytes.
//
// Note that parameter <srcPaths> supports multiple paths join with ','.
func PackToGoFile(srcPath, goFilePath, pkgName string, keyPrefix ...string) error {
	data, err := Pack(srcPath, keyPrefix...)
	if err != nil {
		return err
	}
	return gfile.PutContents(
		goFilePath,
		fmt.Sprintf(gPACKAGE_TEMPLATE, pkgName, bytesToHexStr(data)),
	)
}

// Unpack unpacks the content specified by <path> to []*File.
func Unpack(path string) ([]*File, error) {
	realPath, err := gfile.Search(path)
	if err != nil {
		return nil, err
	}
	return UnpackContent(gfile.GetContents(realPath))
}

// UnpackContent unpacks the content to []*File.
func UnpackContent(content string) ([]*File, error) {
	data, err := gcompress.UnGzip(hexStrToBytes(content))
	if err != nil {
		return nil, err
	}
	reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	array := make([]*File, len(reader.File))
	for i, file := range reader.File {
		array[i] = &File{file: file}
	}
	return array, nil
}

// bytesToHexString converts binary content to hex string content.
func bytesToHexStr(b []byte) string {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	return gconv.UnsafeBytesToStr(dst)
}

// hexStrToBytes converts hex string content to []byte.
func hexStrToBytes(s string) []byte {
	src := gconv.UnsafeStrToBytes(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)
	return dst
}

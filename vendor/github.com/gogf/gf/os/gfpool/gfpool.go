// Copyright 2017-2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gfpool provides io-reusable pool for file pointer.
package gfpool

import (
	"fmt"
	"os"
	"sync"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/container/gpool"
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/os/gfsnotify"
)

// File pointer pool.
type Pool struct {
	id     *gtype.Int  // 指针池ID，用以识别指针池是否需要重建
	pool   *gpool.Pool // 底层对象池
	inited *gtype.Bool // 是否初始化(在执行第一次执行File方法后初始化，主要用于文件监听的添加，但是只能添加一次)
	expire int         // 过期时间
}

// 文件指针池指针
type File struct {
	*os.File              // 底层文件指针
	mu       sync.RWMutex // 互斥锁
	pool     *Pool        // 所属池
	poolid   int          // 所属池ID，如果池ID不同表示池已经重建，那么该文件指针也应当销毁，不能重新丢到原有的池中
	flag     int          // 打开标志
	perm     os.FileMode  // 打开权限
	path     string       // 绝对路径
}

var (
	// 全局文件指针池Map, 不过期
	pools = gmap.NewStrAnyMap(true)
)

// 获得文件对象，并自动创建指针池(过期时间单位：毫秒)
func Open(path string, flag int, perm os.FileMode, expire ...int) (file *File, err error) {
	fpExpire := 0
	if len(expire) > 0 {
		fpExpire = expire[0]
	}
	pool := pools.GetOrSetFuncLock(fmt.Sprintf("%s&%d&%d&%d", path, flag, expire, perm), func() interface{} {
		return New(path, flag, perm, fpExpire)
	}).(*Pool)

	return pool.File()
}

// 创建一个文件指针池，expire = 0表示不过期，expire < 0表示使用完立即回收，expire > 0表示超时回收，默认值为0表示不过期。
// 注意过期时间单位为：毫秒。
func New(path string, flag int, perm os.FileMode, expire ...int) *Pool {
	fpExpire := 0
	if len(expire) > 0 {
		fpExpire = expire[0]
	}
	p := &Pool{
		id:     gtype.NewInt(),
		expire: fpExpire,
		inited: gtype.NewBool(),
	}
	p.pool = newFilePool(p, path, flag, perm, fpExpire)
	return p
}

// 创建文件指针池
func newFilePool(p *Pool, path string, flag int, perm os.FileMode, expire int) *gpool.Pool {
	pool := gpool.New(expire, func() (interface{}, error) {
		file, err := os.OpenFile(path, flag, perm)
		if err != nil {
			return nil, err
		}
		return &File{
			File:   file,
			pool:   p,
			poolid: p.id.Val(),
			flag:   flag,
			perm:   perm,
			path:   path,
		}, nil
	}, func(i interface{}) {
		_ = i.(*File).File.Close()
	})
	return pool
}

// 获得一个文件打开指针
func (p *Pool) File() (*File, error) {
	if v, err := p.pool.Get(); err != nil {
		return nil, err
	} else {
		f := v.(*File)
		stat, err := os.Stat(f.path)
		if f.flag&os.O_CREATE > 0 {
			if os.IsNotExist(err) {
				if file, err := os.OpenFile(f.path, f.flag, f.perm); err != nil {
					return nil, err
				} else {
					f.File = file
					if stat, err = f.Stat(); err != nil {
						return nil, err
					}
				}
			}
		}
		if f.flag&os.O_TRUNC > 0 {
			if stat.Size() > 0 {
				if err := f.Truncate(0); err != nil {
					return nil, err
				}
			}
		}
		if f.flag&os.O_APPEND > 0 {
			if _, err := f.Seek(0, 2); err != nil {
				return nil, err
			}
		} else {
			if _, err := f.Seek(0, 0); err != nil {
				return nil, err
			}
		}
		// 优先使用 !p.inited.Val() 原子读取操作判断，保证判断操作的效率；
		// p.inited.Set(true) == false 使用原子写入操作，保证该操作的原子性；
		if !p.inited.Val() && p.inited.Set(true) == false {
			_, _ = gfsnotify.Add(f.path, func(event *gfsnotify.Event) {
				// 如果文件被删除或者重命名，立即重建指针池
				if event.IsRemove() || event.IsRename() {
					// 原有的指针都不要了
					p.id.Add(1)
					// Clear相当于重建指针池
					p.pool.Clear()
					// 为保证原子操作，但又不想加锁，
					// 这里再执行一次原子Add，将在两次Add中间可能分配出去的文件指针丢弃掉
					p.id.Add(1)
				}
			}, false)
		}
		return f, nil
	}
}

// 关闭指针池
func (p *Pool) Close() {
	p.pool.Close()
}

// 获得底层文件指针(返回error是标准库io.ReadWriteCloser接口实现)
func (f *File) Close() error {
	if f.poolid == f.pool.id.Val() {
		f.pool.pool.Put(f)
	}
	return nil
}

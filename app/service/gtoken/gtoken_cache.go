package gtoken

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
	"time"
)

//token内容
type TokenContent struct {
	UserKey     string      //用户key
	Uuid        string      //随机字符串
	Token       string      //生成的token
	Data        interface{} //数据
	CreateTime  int         //创建时间
	RefreshTime int         //刷新时间
}

// setCache 设置缓存
func (m *GfToken) setCache(cacheKey string, content *TokenContent) bool {
	switch m.CacheMode {
	case CacheModeCache:
		gcache.Set(cacheKey, content, gconv.Duration(m.Timeout)*time.Millisecond)
	case CacheModeRedis:
		cacheValueJson, err1 := gjson.Encode(content)
		if err1 != nil {
			return false
		}
		_, err := g.Redis().Do("SETEX", cacheKey, m.Timeout, cacheValueJson)
		if err != nil {
			return false
		}
	default:
		return false
	}

	return true
}

// getCache 获取缓存
func (m *GfToken) getCache(cacheKey string) *TokenContent {
	var userCache TokenContent
	switch m.CacheMode {
	case CacheModeCache:
		userCacheValue := gcache.Get(cacheKey)
		if userCacheValue == nil {
			return nil
		}
		gconv.Struct(userCacheValue, &userCacheValue)
	case CacheModeRedis:
		userCacheJson, err := g.Redis().Do("GET", cacheKey)
		if err != nil {
			return nil
		}
		if userCacheJson == nil {
			return nil
		}

		err = gjson.DecodeTo(userCacheJson, &userCache)
		if err != nil {
			return nil
		}
	default:
		return nil
	}

	return &userCache
}

// removeCache 删除缓存
func (m *GfToken) removeCache(cacheKey string) bool {
	switch m.CacheMode {
	case CacheModeCache:
		gcache.Remove(cacheKey)
	case CacheModeRedis:
		var err error
		_, err = g.Redis().Do("DEL", cacheKey)
		if err != nil {
			return false
		}
	default:
		return false
	}

	return true
}

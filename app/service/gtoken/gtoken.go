package gtoken

import (
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"strings"
	"yj-app/app/model"
)

const (
	CacheModeCache = 1
	CacheModeRedis = 2
)

var instance *GfToken

// GfToken gtoken结构体
type GfToken struct {
	// 缓存模式 1 gcache 2 gredis 默认1
	CacheMode int8
	// 缓存key
	CacheKey string
	// 超时时间 默认10天
	Timeout int
	// 缓存刷新时间 默认为超时时间的一半
	MaxRefresh int
	// Token分隔符
	TokenDelimiter string
	// Token加密key
	EncryptKey []byte
	// 是否支持多端登录，默认false
	MultiLogin bool
}

// 获取token实例
func Instance() *GfToken {
	if instance != nil {
		return instance
	}

	var token GfToken

	//存储类别
	cacheMode := g.Cfg().GetInt8("api.jwt.CacheMode")
	if cacheMode > 0 {
		token.CacheMode = cacheMode
	} else {
		token.CacheMode = CacheModeCache
	}
	cacheKey := g.Cfg().GetString("api.jwt.CacheKey")
	if cacheKey != "" {
		token.CacheKey = cacheKey
	} else {
		token.CacheKey = "GToken:"
	}

	timeout := g.Cfg().GetInt("api.jwt.Timeout")
	if cacheKey != "" {
		token.Timeout = timeout
	} else {
		token.Timeout = 3600000
	}

	token.MaxRefresh = token.Timeout / 2

	tokenDelimiter := g.Cfg().GetString("api.jwt.TokenDelimiter")
	if tokenDelimiter != "" {
		token.TokenDelimiter = tokenDelimiter
	} else {
		token.TokenDelimiter = "_"
	}

	encryptKey := g.Cfg().GetString("api.jwt.EncryptKey")
	if encryptKey != "" {
		token.EncryptKey = []byte(encryptKey)
	} else {
		token.EncryptKey = []byte("yjgoyjgoyjgoyjgo")
	}

	multiLogin := g.Cfg().GetBool("api.jwt.MultiLogin")
	token.MultiLogin = multiLogin

	instance = &token
	return instance
}

// Login 登录
func (m *GfToken) Login(userKey string, data interface{}) *TokenContent {
	if m.MultiLogin {
		// 支持多端重复登录，获取已经存在的token
		tc := m.getToken(userKey)
		if tc != nil {
			tmp, _ := m.EncryptToken(userKey, tc.Uuid)
			return tmp
		}
	}
	// 生成token
	tc, _ := m.genToken(userKey, data)
	return tc
}

// Logout 登出
func (m *GfToken) Logout(r *ghttp.Request) {
	// 获取请求token
	token := m.getRequestToken(r)
	if token != "" {
		// 删除token
		m.removeToken(token)
	}
}

// GetTokenData 通过token获取对象
func (m *GfToken) GetTokenData(r *ghttp.Request) model.CommonRes {
	token := m.getRequestToken(r)
	if token != "" {
		// 验证token
		err := m.validToken(token)
		if err == nil {
			return model.CommonRes{
				Code:  model.SUCCESS,
				Msg:   "验证成功",
				Data:  token,
				Btype: model.Buniss_Other,
			}
		} else {
			return model.CommonRes{
				Code:  model.FAIL,
				Msg:   err.Error(),
				Data:  "",
				Btype: model.Buniss_Other,
			}
		}
	}

	return model.CommonRes{
		Code:  model.FAIL,
		Msg:   "获取Token失败",
		Data:  "",
		Btype: model.Buniss_Other,
	}
}

// getRequestToken 返回请求Token
func (m *GfToken) getRequestToken(r *ghttp.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			return ""
		} else if parts[1] == "" {
			return ""
		}

		return parts[1]
	}

	authHeader = r.Header.Get("token")
	return authHeader

}

// genToken 生成Token
func (m *GfToken) genToken(userKey string, data interface{}) (*TokenContent, error) {
	token, err := m.EncryptToken(userKey, "")
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, gerror.New("生成Token失败")
	}

	cacheKey := m.CacheKey + userKey

	token.Data = data
	token.CreateTime = gtime.Now().Millisecond()
	token.RefreshTime = gtime.Now().Millisecond()

	rs := m.setCache(cacheKey, token)
	if !rs {
		return nil, gerror.New("保存token失败")
	}
	return token, nil
}

// validToken 验证Token
func (m *GfToken) validToken(token string) error {
	if token == "" {
		return gerror.New("参数错误")
	}

	decryptToken := m.DecryptToken(token)
	if decryptToken == nil {
		return gerror.New("token解密错误")
	}

	gt := m.getToken(decryptToken.UserKey)
	if gt == nil {
		return gerror.New("token获取失败")
	}

	if decryptToken.Uuid != gt.Uuid {
		return gerror.New("token验证错误")
	}

	return nil
}

// getToken 通过userKey获取Token
func (m *GfToken) getToken(userKey string) *TokenContent {
	cacheKey := m.CacheKey + userKey

	cache := m.getCache(cacheKey)
	if cache == nil {
		return nil
	}

	nowTime := gtime.Now().Millisecond()
	refreshTime := cache.RefreshTime

	// 需要进行缓存超时时间刷新
	if gconv.Int64(refreshTime) == 0 || nowTime > gconv.Int(refreshTime) {
		cache.CreateTime = gtime.Now().Millisecond()
		cache.RefreshTime = gtime.Now().Millisecond() + m.MaxRefresh
		m.setCache(cacheKey, cache)
	}

	return cache
}

// removeToken 删除Token
func (m *GfToken) removeToken(token string) bool {
	decryptToken := m.DecryptToken(token)
	cacheKey := m.CacheKey + decryptToken.UserKey
	return m.removeCache(cacheKey)
}

// EncryptToken token加密方法
func (m *GfToken) EncryptToken(userKey string, uuid string) (*TokenContent, error) {
	if userKey == "" {
		return nil, gerror.New("用户密钥不能为空")
	}

	var t TokenContent

	if uuid == "" {
		// 重新生成uuid
		newUuid, err := gmd5.Encrypt(grand.Str(10))
		if err != nil {
			return nil, err
		}
		uuid = newUuid
	}

	tokenStr := userKey + m.TokenDelimiter + uuid
	tmp := []byte(tokenStr)
	token, err := gaes.Encrypt(tmp, m.EncryptKey)
	if err != nil {
		return nil, err
	}
	t.UserKey = userKey
	t.Uuid = uuid
	t.Token = string(gbase64.Encode(token))

	return &t, nil
}

// DecryptToken token解密方法
func (m *GfToken) DecryptToken(token string) *TokenContent {
	var t TokenContent
	if token == "" {
		return nil
	}

	token64, err := gbase64.Decode([]byte(token))
	if err != nil {
		return nil
	}
	decryptToken, err2 := gaes.Decrypt([]byte(token64), m.EncryptKey)
	if err2 != nil {
		return nil
	}
	tokenArray := gstr.Split(string(decryptToken), m.TokenDelimiter)
	if len(tokenArray) < 2 {
		return nil
	}

	t.UserKey = tokenArray[0]
	t.Uuid = tokenArray[1]
	t.Token = token

	return &t
}

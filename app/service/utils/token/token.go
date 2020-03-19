package token

import (
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"sync"
	"time"
)

var (
	instance *YToken
	once     sync.Once
)

// YToken gtoken结构体
type YToken struct {
	// 超时时间 单位秒
	Timeout int64
	// token更新时间 默认为超时时间的一半
	MaxRefresh int64
	// Token分隔符
	TokenDelimiter string
	// Token加密key
	EncryptKey []byte
}

//token内容
type TokenContent struct {
	UserKey     string //用户key
	Token       string //生成的token
	NewToken    string //如果过了刷新时间生成新的token
	Data        string //数据
	RefreshTime int64  //刷新时间
	ExpireTime  int64  // 过期时间
}

// 获取token实例
func Instance() *YToken {
	once.Do(func() {
		var token YToken
		timeout := g.Cfg().GetInt64("api.jwt.Timeout")
		if timeout > 0 {
			token.Timeout = timeout
		} else {
			token.Timeout = 3600
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

		instance = &token
	})

	return instance
}

// Login 登录
func (m *YToken) Login(userKey string, data string) *TokenContent {
	// 生成token
	tc, _ := m.EncryptToken(userKey, data)
	return tc
}

// GetTokenData 通过token获取对象
func (m *YToken) GetTokenData(r *ghttp.Request) (*TokenContent, error) {
	token := m.getRequestToken(r)
	if token != "" {
		// 验证token
		return m.validToken(token)
	}
	return nil, gerror.New("获取token错误")
}

// getRequestToken 返回请求Token
func (m *YToken) getRequestToken(r *ghttp.Request) string {
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

// validToken 验证Token
func (m *YToken) validToken(tokenStr string) (*TokenContent, error) {
	if tokenStr == "" {
		return nil, gerror.New("参数错误")
	}

	token := m.DecryptToken(tokenStr)
	if token == nil {
		return nil, gerror.New("token解密错误")
	}
	//判断是否过期
	if token.ExpireTime < time.Now().Unix() {
		return nil, gerror.New("token已过期")
	}

	//判断是否需要刷新
	if token.RefreshTime > time.Now().Unix() {
		//刷新token
		newToken, _ := m.EncryptToken(token.UserKey, token.Data)
		token.NewToken = newToken.Token
	}

	return token, nil
}

// EncryptToken token加密方法  userKey：  用户标识字符串 data：  用户数据字符串
func (m *YToken) EncryptToken(userKey string, data string) (*TokenContent, error) {
	if userKey == "" {
		return nil, gerror.New("用户密钥不能为空")
	}

	var t TokenContent

	t.Data = data
	t.RefreshTime = time.Now().Unix() + m.MaxRefresh
	t.ExpireTime = time.Now().Unix() + m.Timeout
	t.UserKey = userKey

	//用户标识字符串 分隔符 数据 分隔符 刷新时间 分隔符 过期时间
	tokenStr := userKey + m.TokenDelimiter + data + m.TokenDelimiter + gconv.String(t.RefreshTime) + m.TokenDelimiter + gconv.String(t.ExpireTime)
	tmp := []byte(tokenStr)
	tokenArr, err := gaes.Encrypt(tmp, m.EncryptKey)
	if err != nil {
		return nil, err
	}
	tokenArr64 := gbase64.Encode(tokenArr)
	t.Token = string(tokenArr64)
	return &t, nil
}

// DecryptToken token解密方法
func (m *YToken) DecryptToken(tokenStr string) *TokenContent {
	var token TokenContent
	if tokenStr == "" {
		return nil
	}

	tmp64, err := gbase64.Decode([]byte(tokenStr))

	if err != nil {
		return nil
	}

	decryptToken, err := gaes.Decrypt(tmp64, m.EncryptKey)
	if err != nil {
		return nil
	}
	tokenArray := gstr.Split(string(decryptToken), m.TokenDelimiter)
	if len(tokenArray) != 4 {
		return nil
	}
	//0:用户标识字符串  1:数据  2:刷新时间  3:过期时间
	token.UserKey = tokenArray[0]
	token.Data = tokenArray[1]
	token.RefreshTime = gconv.Int64(tokenArray[2])
	token.ExpireTime = gconv.Int64(tokenArray[3])
	token.Token = tokenStr

	return &token
}

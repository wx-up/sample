package jwt

import (
	"errors"
	"strings"
	"time"

	"sample/pkg/app"

	"github.com/gin-gonic/gin"

	"sample/pkg/logger"

	jwtPkg "github.com/golang-jwt/jwt"
	"sample/pkg/config"
)

var (
	ErrTokenExpired           = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         = errors.New("请求令牌格式有误")
	ErrTokenInvalid           = errors.New("请求令牌无效")
	ErrHeaderEmpty            = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        = errors.New("请求头中 Authorization 格式有误")
)

type jwt struct {
	key        []byte
	maxRefresh time.Duration
	expireTime int64
}

type Claims struct {
	Uid      string `json:"uid"`
	UserName string `json:"user_name"`
	jwtPkg.StandardClaims
}

type Option func(j *jwt)

func WithKey(key string) Option {
	return func(j *jwt) {
		j.key = []byte(key)
	}
}

func WithMaxRefresh(t time.Duration) Option {
	return func(j *jwt) {
		j.maxRefresh = t
	}
}

func WithExpireTime(expireTime int64) Option {
	return func(j *jwt) {
		j.expireTime = expireTime
	}
}

func New(opts ...Option) *jwt {
	res := &jwt{
		key:        []byte(config.Get("app.key")),
		maxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

// ParseToken 解析token
func (j *jwt) ParseToken(ctx *gin.Context) (*Claims, error) {
	tokenString, err := j.getTokenFromHeader(ctx)
	if err != nil {
		return nil, err
	}

	return j.ParseTokenBy(tokenString)
}

func (j *jwt) ParseTokenBy(tokenString string) (*Claims, error) {
	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := j.parseToken(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		if !ok {
			return nil, ErrTokenInvalid
		}
		err := ErrTokenMalformed
		if validationErr.Errors == jwtPkg.ValidationErrorExpired {
			err = ErrTokenExpired
		}
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

func (j *jwt) RefreshToken(ctx *gin.Context) (string, error) {
	// 1. 从 Header 里获取 token
	tokenString, err := j.getTokenFromHeader(ctx)
	if err != nil {
		return "", err
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := j.parseToken(tokenString)
	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwtPkg.ValidationError)
		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwtPkg.ValidationErrorExpired {
			return "", err
		}
	}

	// 4. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*Claims)

	// 5. 检查是否过了『最大允许刷新的时间』
	unix := app.TimeNowInTimezone().Add(j.maxRefresh * -1).Unix()
	if claims.IssuedAt <= unix {
		return "", ErrTokenExpiredMaxRefresh
	}
	// 修改过期时间
	timeNow := j.getExpireTime()
	claims.StandardClaims.ExpiresAt = timeNow
	return j.generateToken(claims)
}

func (j *jwt) parseToken(token string) (*jwtPkg.Token, error) {
	return jwtPkg.ParseWithClaims(token, &Claims{}, func(token *jwtPkg.Token) (interface{}, error) {
		return j.key, nil
	})
}

func (j *jwt) getTokenFromHeader(ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

// GenerateToken 创建 token
func (j *jwt) GenerateToken(uid string, userName string) string {
	expireTime := j.getExpireTime()
	timeNow := app.TimeNowInTimezone().Unix()
	claims := &Claims{
		Uid:      uid,
		UserName: userName,
		StandardClaims: jwtPkg.StandardClaims{
			ExpiresAt: expireTime, // 过期时间
			IssuedAt:  timeNow,    // 首次签名时间
			Issuer:    config.Get("app.name"),
			NotBefore: timeNow, // 生效时间
		},
	}
	token, err := j.generateToken(claims)
	logger.LogIf(err)
	return token
}

// generateToken 生成 Token
func (j *jwt) generateToken(claims *Claims) (string, error) {
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, claims)
	return token.SignedString(j.key)
}

// getExpireTime 获取过期时间：时间戳
func (j *jwt) getExpireTime() int64 {
	if j.expireTime > 0 {
		return j.expireTime
	}
	timeNow := app.TimeNowInTimezone()

	expireTime := config.GetInt64("jwt.debug_expire_time")
	if !config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	return timeNow.Add(time.Duration(expireTime) * time.Minute).Unix()
}

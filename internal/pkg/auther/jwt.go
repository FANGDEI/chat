package auther

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	jardiniere "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func init() {
	defaultTokenerManager = New()
}

var defaultTokenerManager *Manager

type Manager struct {
	tokener *jardiniere.Middleware
}

func New() *Manager {
	return &Manager{
		tokener: jardiniere.New(jardiniere.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("chat"), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
			// 自定义token提取器
			Extractor: func(ctx iris.Context) (string, error) {
				authHeader := ctx.GetHeader("Authorization")
				if authHeader == "" {
					if authHeader = ctx.GetHeader("Sec-WebSocket-Protocol"); authHeader == "" {
						return "", nil // No error, just no token
					}
				}
				authHeaderParts := strings.Split(authHeader, " ")
				if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
					return "", fmt.Errorf("authorization header format must be Bearer {token}")
				}
				return authHeaderParts[1], nil
			},
			ErrorHandler: func(ctx iris.Context, err error) {
				ctx.StatusCode(iris.StatusForbidden)
				log.Printf("[%s] CODE(%d) ERROR : %+v\n", ctx.Path(), iris.StatusForbidden, err)
				ctx.JSON(map[string]any{
					"msg": "权限认证失败",
				})
			},
		}),
	}
}

func (m *Manager) NewToken(id int64) (tokenString string, err error) {
	t := time.Now()
	message := jwt.MapClaims{
		"iat": t.Unix(),
		"exp": t.Add(7 * 24 * time.Hour).Unix(),
		"iss": "CHAT",
		"id":  id,
	}
	tokenString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, message).SignedString([]byte("chat"))
	return
}

func (m *Manager) Serve() func(ctx iris.Context) {
	return m.tokener.Serve
}

func (m *Manager) GetID(ctx iris.Context) int64 {
	return int64(ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(float64))
}

//func (m *Manager) GetUUID(ctx iris.Context) string {
//	return ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)["uuid"].(string)
//}

func GetDefaultTokenerManager() *Manager {
	return defaultTokenerManager
}

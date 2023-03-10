package auther

import (
	"github.com/golang-jwt/jwt/v4"
	jardiniere "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"log"
	"time"
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

func (m *Manager) NewToken(id int64, uuid string) (tokenString string, err error) {
	t := time.Now()
	message := jwt.MapClaims{
		"iat":  t.Unix(),
		"exp":  t.Add(7 * 24 * time.Hour).Unix(),
		"iss":  "CHAT",
		"id":   id,
		"uuid": uuid,
	}
	tokenString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, message).SignedString([]byte("chat"))
	return
}

func (m *Manager) getID(ctx iris.Context) int64 {
	return int64(ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)["id"].(float64))
}

func (m *Manager) getUUID(ctx iris.Context) string {
	return ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)["uuid"].(string)
}

func GetDefaultTokenerManager() *Manager {
	return defaultTokenerManager
}

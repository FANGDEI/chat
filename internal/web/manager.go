package web

import (
	"chat/internal/pkg/auther"
	"chat/internal/pkg/localer"
	"chat/internal/pkg/logger"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"log"
	"reflect"
	"strings"
)

var defaultLogger = logger.GetDefaultLoggerManager()

type Manager struct {
	handler *iris.Application
	logger  *logger.Manager
	tokener *auther.Manager
	localer *localer.Manager
}

func New() *Manager {
	return &Manager{
		handler: iris.Default(),
		logger:  logger.GetDefaultLoggerManager(),
		tokener: auther.GetDefaultTokenerManager(),
		localer: localer.GetDefaultLocalerManager(),
	}
}

func (m *Manager) Run() error {
	err := m.load()
	if err != nil {
		return err
	}
	return m.handler.Run(iris.Addr(C.Host + ":" + C.Port))
}

func (m *Manager) load() (err error) {
	err = m.loadPlugin()
	if err != nil {
		return
	}
	return m.loadRoute()
}

func (m *Manager) loadPlugin() error {
	m.loadCORS()
	return nil
}

func (m *Manager) loadCORS() error {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "DELETE"},
		MaxAge:           3600,
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	m.handler.UseRouter(crs)
	return nil
}

func (m *Manager) loadRoute() error {
	t := reflect.TypeOf(m)
	for i := 0; i < t.NumMethod(); i++ {
		f := t.Method(i)
		if strings.HasPrefix(f.Name, "Route") &&
			f.Type.NumOut() == 0 &&
			f.Type.NumIn() == 1 {
			log.Println("[GATEWAY] LOAD ROUTE:", f.Name)
			f.Func.Call([]reflect.Value{
				reflect.ValueOf(m),
			})
		}
	}
	return nil
}

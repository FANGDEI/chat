package localer

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	var err error
	defaultLocalerManager, err = New()
	if err != nil {
		log.Println(err)
	}
}

var defaultLocalerManager *Manager

type Manager struct {
	handler *gorm.DB
}

func New() (*Manager, error) {
	db, err := gorm.Open(
		mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				C.User,
				C.Password,
				C.Host,
				C.Port,
				C.Name,
			),
		),
		&gorm.Config{},
	)
	return &Manager{
		handler: db,
	}, err
}

func (m *Manager) execTx(fn func(*Manager) error) error {
	tx := m.handler.Begin()
	manager := &Manager{handler: tx}
	err := fn(manager)
	if err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit().Error
}

func GetDefaultLocalerManager() *Manager {
	return defaultLocalerManager
}

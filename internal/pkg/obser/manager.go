package obser

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var defaultObserManager *Manager

func init() {
	var err error
	defaultObserManager, err = New()
	if err != nil {
		log.Fatalln(err)
	}
}

type Manager struct {
	handler *minio.Client
}

func New() (*Manager, error) {
	handler, err := minio.New(C.Address, &minio.Options{
		Creds: credentials.NewStaticV4(
			C.SecretID,
			C.SecretKey,
			"",
		),
	})
	return &Manager{
		handler: handler,
	}, err
}

func GetDefaultObserManager() *Manager {
	return defaultObserManager
}

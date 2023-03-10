package cryptoer

import "log"

func init() {
	var err error
	defaultCryptoerManager, err = New()
	if err != nil {
		log.Println(err)
	}
}

var defaultCryptoerManager *Manager

type Manager struct {
}

func New() (*Manager, error) {
	return &Manager{}, nil
}

func GetDefaultCryptoerManager() *Manager {
	return defaultCryptoerManager
}

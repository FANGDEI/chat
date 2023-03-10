package chat

import (
	"chat/internal/app/service"
	"chat/internal/pkg/auther"
	"chat/internal/pkg/cacher"
	"chat/internal/pkg/cryptoer"
	"chat/internal/pkg/emailer"
	"chat/internal/pkg/localer"
	"chat/internal/pkg/logger"
)

type Manager struct {
	service.UnimplementedChatServiceServer
	localer  *localer.Manager
	cacher   *cacher.Manager
	logger   *logger.Manager
	cryptoer *cryptoer.Manager
	emailer  *emailer.Manager
	tokener  *auther.Manager
}

func New() *Manager {
	return &Manager{
		localer:  localer.GetDefaultLocalerManager(),
		cacher:   cacher.GetDefaultCacherManager(),
		logger:   logger.GetDefaultLoggerManager(),
		cryptoer: cryptoer.GetDefaultCryptoerManager(),
		emailer:  emailer.GetDefaultEmailerManager(),
		tokener:  auther.GetDefaultTokenerManager(),
	}
}

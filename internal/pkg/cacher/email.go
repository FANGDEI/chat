package cacher

import (
	"context"
	"errors"
	"time"
)

func (m *Manager) SetEmailCode(email, code string) error {
	return m.handler.Set(context.Background(), m.getEmailVerifyKey(email), code, time.Minute*15).Err()
}

func (m *Manager) GetEmailCode(email string) (string, error) {
	return m.handler.Get(context.Background(), m.getEmailVerifyKey(email)).Result()
}

func (m *Manager) DelEmailCode(email string) error {
	return m.handler.Del(context.Background(), m.getEmailVerifyKey(email)).Err()
}

func (m *Manager) AuthEmail(email, code string) error {
	if str, err := m.GetEmailCode(email); err != nil {
		return err
	} else if str != code {
		return errors.New("code is not true")
	} else {
		return m.DelEmailCode(email)
	}
}

func (m *Manager) SetEmailToBan(email string) error {
	return m.handler.Set(context.Background(), m.getEmailBanKey(email), 1, time.Minute).Err()
}

func (m *Manager) IsEmailBanned(email string) bool {
	num, err := m.handler.Exists(context.Background(), m.getEmailBanKey(email)).Result()
	return err == nil && num != 0
}

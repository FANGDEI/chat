package cacher

import (
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	defaultCacherManager.Send(&Message{
		From:        1,
		To:          2,
		Content:     "test3",
		ContentType: 1,
		MessageType: 1,
		Time:        "1",
	})
}

func TestGetMsg(t *testing.T) {
	s, err := defaultCacherManager.GetMsg(2, time.Hour)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%v\n", s[0])
	t.Log(s)
}

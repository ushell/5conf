package log

import (
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type hookLog struct {
	TraceId string
}

func newHookLog() logrus.Hook {
	instance := new(hookLog)
	return instance
}

func (h *hookLog) Levels() [] logrus.Level {
	return logrus.AllLevels
}

func (h *hookLog) Fire(entry *logrus.Entry) error {
	h.TraceId = generateUUID()
	entry.Data["trace_id"] = h.TraceId
	return nil
}

func (h *hookLog) GetTraceId() string {
	return h.TraceId
}

func generateUUID() string {
	return uuid.NewV1().String()
}

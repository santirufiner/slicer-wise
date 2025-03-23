package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type logHook struct{}

func newLogHook() *logHook {
	return &logHook{}
}

func (hook logHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook logHook) Fire(entry *logrus.Entry) error {
	// Skip 7 frames to get to the caller of the function that called logrus
	// If there is additional fields, skip 6 frames to get the right caller
	skip := 7
	if len(entry.Data) > 0 {
		skip = 6
	}

	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return nil
	}

	// Split the file path at "slicer-wise" and take the second part
	parts := strings.Split(file, "slicer-wise")
	if len(parts) > 1 {
		file = parts[1]
	}
	entry.Data["caller"] = fmt.Sprintf("%s:%d", file, line)

	return nil
}

package logger

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const ISO8601 = "2006-02-03 15:04:05"

type SystemLog struct {
	Level   string `json:"level"`
	Date    string `json:"time"`
	Queue   string `json:"queue"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

func CreateMessageLog(err error) error {
	pc, fn, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("in function:unknown file:unknown line:unknown message: %v", err)
	}
	nameFull := runtime.FuncForPC(pc).Name()
	nameEnd := filepath.Ext(nameFull)
	name := strings.TrimPrefix(nameEnd, ".")
	if name == "0" {
		name = "init"
	}
	return fmt.Errorf("in function:%s file:%s line:%d message:%v", name, filepath.Base(fn), line, err)
}

func PrintLog(level, queue, id string, msg interface{}) {
	systemLog := NewSystemLog(
		level,
		fmt.Sprintf("%s", msg),
		queue,
		id,
	)
	fmt.Println(LogToJson(systemLog))
}

func NewSystemLog(level, message, queue, id string) *SystemLog {
	t := time.Now()
	return &SystemLog{
		Level:   level,
		Date:    t.Format(ISO8601),
		Queue:   queue,
		Message: message,
		ID:      id,
	}
}

func LogToJson(log interface{}) string {
	data, err := json.Marshal(log)
	if err != nil {
		fmt.Println("[ERROR]: Failed to logging ", err.Error())
		return ""
	}
	return string(data)
}

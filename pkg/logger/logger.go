package logger

import (
	"errors"
	"fmt"
	_ "github.com/mailru/easyjson/gen"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const ISO8601 = "2006-02-03 15:04:05"

const (
	LevelError  = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

//easyjson
type logger struct {
	level   uint8    `json:"omitempty"`
	Level   string `json:"level"`
	Date    string `json:"time"`
	Queue   string `json:"queue"`
	ID      string `json:"id"`
	Message string `json:"message"`
}

func (l *logger) SetLevel(level string) error {
	switch level{
	case "DEBUG":
		l.level=LevelDebug
	case "ERROR":
		l.level=LevelError
	case "WARN":
		l.level=LevelWarning
	case "INFO":
		l.level=LevelInformational
	default:
		return errors.New("level undefined")
	}
	return nil
}

func NewLogger(Queue string) *logger {
	return &logger{
		Queue: Queue,
	}
}

func (l *logger) SetID(ID string){
	l.ID=ID
}

func (l *logger) Info(msg string) {
	l.Message = msg
	l.Level = "INFO"
	l.Date = time.Now().Format(ISO8601)
	fmt.Println(l.toJson())
}

func (l *logger) Warn(msg string) {
	if l.level<=LevelWarning{
		return
	}
	l.Level="WARN"
	l.Message = msg
	fmt.Println(l.toJson())
}

func (l *logger) Error(err error) {
	msg := createMessageLog(err)
	l.Message = msg.Error()
	l.Level="WARN"
	l.Date = time.Now().Format(ISO8601)
	fmt.Println(l.toJson())
}

func (l *logger) Debug(msg string) {
	if l.level<LevelError{
		return
	}
	l.Level="DEBUG"
	l.Message = msg
	l.Date = time.Now().Format(ISO8601)
	fmt.Println(l.toJson())
}

func (l *logger) toJson() string {
	data, err := l.MarshalJSON()
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func createMessageLog(err error) error {
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

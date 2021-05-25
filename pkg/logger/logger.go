package logs

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	RFC850  = "Monday, 02-Jan-06 15:04:05 MST"
	ISO8601 = "2006-01-02 15:04:05"
	ISO8602 = "2006-01-02T15:04:05"
)

type StandardLog struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Message string `json:"msg"`
	AppName string `json:"app_name"`
	LogType string `json:"log_type"`
	ID      string `json:"id"`
}

type RequestLog struct {
	Level     string `json:"level"`
	URL       string `json:"request_url"`
	Addr      string `json:"remote_addr"`
	Date      string `json:"time"`
	AppName   string `json:"app_name"`
	Code      *int   `json:"status_code,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
	Message   string `json:"message"`
	LogType   string `json:"log_type"`
	ID        string `json:"id"`
}

func GetRequestLog(lvl, message, serviceID, URL, remoteAddr, userAgent, logType, ID string, code int) *RequestLog {
	t := time.Now()
	return &RequestLog{
		Level:     lvl,
		URL:       URL,
		Addr:      remoteAddr,
		Date:      t.Format(ISO8601),
		AppName:   serviceID,
		Code:      &code,
		UserAgent: userAgent,
		Message:   message,
		LogType:   logType,
		ID:        ID,
	}
}

func GetStandardLog(lvl, msg, appName, logType, ID string) *StandardLog {
	t := time.Now()
	return &StandardLog{
		Level:   lvl,
		Time:    t.Format(ISO8601),
		Message: msg,
		AppName: appName,
		LogType: logType,
		ID:      ID,
	}
}

func Logging(log interface{}) error {
	jsonData, err := json.Marshal(log)
	if err != nil {
		fmt.Println("[ERROR]: Failed to marshal json -", err.Error())
		return err
	} else {
		fmt.Println(jsonData)
		return nil
	}
}

package logapi

import (
	"fmt"
	"time"
)

type LogApi struct {
	DateTime time.Time
	// file log
	Appname  string
	FileName string
	PathName string
	LogData  map[string]interface{}
}

func (l *LogApi) FullFormatLogs() {
	//fmt.Sprintf("{{\"time\":\"%s\"}", "${time}") + ", {\"PID\":\"${pid}\"}, {\"requestID\":\"${locals:requestid}\"}, {\"status_code\":${status}}, {\"method\":\"${method}\"},{\"endpoint\":\"${path}​\"}}\n​",
	log := "{"
	max := len(l.LogData)
	count := 0
	for key, value := range l.LogData {
		count++
		log += fmt.Sprintf("{\"%s\" : %s}", key, value)
		if count != max {
			log += ", "
		}
	}

	log += "}\n"
	fmt.Println(log)
}

func (l *LogApi) Set(key string, value interface{}) {
	if l.LogData == nil {
		l.LogData = make(map[string]interface{})
	}
	l.LogData[key] = value
}

func (l *LogApi) Get(key string) (value interface{}) {
	value = l.LogData[key]
	return
}

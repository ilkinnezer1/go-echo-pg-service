package config

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
)

type RequestLog struct {
	Method      string      `json:"method"`
	Path        string      `json:"path"`
	Headers     http.Header `json:"headers"`
	Body        interface{} `json:"body"`
	QueryParams string      `json:"query_params"`
}

type ResponseLog struct {
	Status int         `json:"status"`
	Header http.Header `json:"header"`
	Body   interface{} `json:"body"`
}

var LoggerConfig = &lumberjack.Logger{
	Filename:   "logs/request.log",
	MaxSize:    100 / 1024, // 100 KB
	MaxBackups: 5,
	MaxAge:     5, // days,
	Compress:   true,
}

package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"main/config"
	"os"
)

// Log MultiWriter
var mw = io.MultiWriter(os.Stdout, config.LoggerConfig)

// HTTPBodyDumpResponse outputs logs for request and response body
func HTTPBodyDumpResponse(c echo.Context, reqBody, resBody []byte) {
	reqLog := config.RequestLog{
		Method:      c.Request().Method,
		Path:        c.Path(),
		Headers:     c.Request().Header,
		QueryParams: c.QueryString(),
	}

	if err := json.Unmarshal(reqBody, &reqLog.Body); err != nil {
		reqLog.Body = string(reqBody)
	}

	resLoq := config.ResponseLog{
		Status: c.Response().Status,
		Header: c.Response().Header(),
	}

	if err := json.Unmarshal(resBody, &resLoq.Body); err != nil {
		resLoq.Body = string(resBody)
	}

	log := map[string]interface{}{
		"Request Body:":  reqLog,
		"Response Body:": resLoq,
	}

	// TODO: Uncomment log Print to see more details about request and response
	_, _ = json.MarshalIndent(log, "", " ")

	//fmt.Println(string(logAsJSON))
}

var LoggerMiddlewareConfig = middleware.LoggerWithConfig(middleware.LoggerConfig{
	Format:           "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, host=${host}, remote_ip=${remote_ip}, user_agent=${user_agent}\n",
	Output:           mw,
	CustomTimeFormat: "2006-01-02 15:04:05",
})

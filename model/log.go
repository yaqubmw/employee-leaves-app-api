package model

import "time"

type RequestLog struct {
	StartTime  time.Time
	EndTime    time.Duration
	StatusCode int
	ClientIP   string
	Method     string
	Path       string
	UserAgent  string
}

// 2023/08/02 - 13:04:07 | 200 |   16.507341ms |       127.0.0.1 | GET      "/api/v1/products" POSTMAN

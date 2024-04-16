package main

import (
	"context"
	"log-service/data"
	"log-service/logs"
)

type LogService struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogService) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	res := &logs.LogResponse{Result: "logeed!"}
	return res, nil
}

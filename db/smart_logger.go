package db

import (
	"fmt"
	"github.com/xaionaro/reform"
	"runtime"
	"strings"
	"time"
)

const (
	SQL_LOGGER_TRACEBACK_DEPTH int = 10
)

type smartLogger struct {
	dbName      string
	traceLogger reform.Logger
	errorLogger reform.Logger
	traceEnable bool
	errorEnable bool
}

func (logger smartLogger) queryWrapper(query string) string {
	var where string

	for i := 2; i < 32; i++ {
		_, filePath, _, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if strings.HasSuffix(filePath, "_reform.go") {
			continue
		}
		if strings.HasSuffix(filePath, "_common.go") {
			continue
		}
		if filePath == `<autogenerated>` {
			continue
		}
		if strings.HasSuffix(filePath, "_generatedModelFunctions.go") {
			continue
		}
		if strings.HasSuffix(filePath, "_generatedMethods.go") {
			continue
		}

		whereArray := make([]string, SQL_LOGGER_TRACEBACK_DEPTH, SQL_LOGGER_TRACEBACK_DEPTH)
		for j := 0; j < SQL_LOGGER_TRACEBACK_DEPTH; j++ {
			_, filePath, line, ok := runtime.Caller(i + j)
			pathParts := strings.Split(filePath, "/")
			fileName := pathParts[len(pathParts)-1]
			if !ok || strings.HasSuffix(fileName, ".s") {
				whereArray = whereArray[SQL_LOGGER_TRACEBACK_DEPTH-j:]
				break
			}
			whereArray[SQL_LOGGER_TRACEBACK_DEPTH-1-j] = fmt.Sprintf("%v:%v", fileName, line)
		}
		where = "[" + strings.Join(whereArray, " -> ") + "] "
		break
	}
	return fmt.Sprintf("%v[db:%s] %s", where, logger.dbName, query)
}
func (logger *smartLogger) SetTraceEnable(enable bool) {
	logger.traceEnable = enable
}
func (logger *smartLogger) SetErrorEnable(enable bool) {
	logger.errorEnable = enable
}
func (logger smartLogger) Before(query string, args []interface{}) {
	if logger.traceEnable {
		logger.traceLogger.Before(logger.queryWrapper(query), args)
	}
	return
}
func (logger smartLogger) After(query string, args []interface{}, d time.Duration, err error) {
	if err != nil {
		if logger.errorEnable {
			logger.errorLogger.After(logger.queryWrapper(query), args, d, err)
		}
	} else {
		if logger.traceEnable {
			logger.traceLogger.After(logger.queryWrapper(query), args, d, err)
		}
	}
	return
}

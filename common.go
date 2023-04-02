package goantlr

import (
	"github.com/araddon/dateparse"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

const (
	MonthsPerYear = 12    // 每年12个月
	SecondsPerDay = 86400 // 每天86400秒

	DateFormatLayer     = "2006-01-02"
	MonthFormatLayer    = "2006-01"
	TimeFormatLayer     = "15:04:05"
	TimeFormatLayer1    = "15:04"
	DateTimeFormatLayer = "2006-01-02 15:04:05"
)

type ResultType int

const (
	ResultError              = -1
	ResultInteger ResultType = iota + 1 //整数
	ResultFloat                         // 浮点数
	ResultString                        // 字符串
	ResultBool                          // 布尔值
)

var (
	codeRegex  = regexp.MustCompile("#(.*?)#")
	codeRegex1 = regexp.MustCompile(`#(.*?)#|"(.*?")"`)
	codeRegex2 = regexp.MustCompile(`"(.*?")"|\(|\)`)
)

func ParseDateTime(s string) (time.Time, error) {
	dt, err := dateparse.ParseIn(s, time.Local) // 常见的日期时间格式 https://github.com/araddon/dateparse
	if err == nil {
		return dt, nil
	}
	// 如果解析没成功，返回无效时间和error
	return time.Time{}, err
}

func ParseResultType(srcVal interface{}) ResultType {
	if srcVal == nil {
		return ResultError
	}
	getType := reflect.TypeOf(srcVal)
	switch getType.Kind() {
	case reflect.String:
		return ResultString
	case reflect.Int, reflect.Int32, reflect.Int64:
		return ResultInteger
	case reflect.Float32, reflect.Float64:
		return ResultFloat
	case reflect.Bool:
		return ResultBool

	default:
		return ResultError
	}
}

func Interface2Int64(srcVal interface{}) (int64, error) {
	if srcVal == nil {
		return 0, nil
	}
	getType := reflect.TypeOf(srcVal)
	switch getType.Kind() {
	case reflect.String:
		{
			intV, err := strconv.ParseInt(srcVal.(string), 10, 64)
			if err != nil {
				return 0, err
			}
			return intV, nil
		}
	case reflect.Int:
		return int64(srcVal.(int)), nil
	case reflect.Int32:
		return int64(srcVal.(int32)), nil
	case reflect.Int64:
		return srcVal.(int64), nil
	case reflect.Float32:
		return int64(srcVal.(float32)), nil
	case reflect.Float64:
		return int64(srcVal.(float64)), nil
	default:
		return 0, nil
	}
}

func ReflectFloat64(val interface{}) (float64, error) {
	if val == nil || val == "" {
		return 0, nil
	}
	switch val.(type) {
	case float32:
		return float64(val.(float32)), nil
	case float64:
		return val.(float64), nil
	case int:
		return float64(val.(int)), nil
	case int32:
		return float64(val.(int32)), nil
	case int64:
		return float64(val.(int64)), nil
	case string:
		v, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	default:
		return 0, nil
	}
}

func IsNumber(val interface{}) bool {
	if val == nil || val == "" {
		return false
	}
	switch val.(type) {
	case float32:
		return true
	case float64:
		return true
	case int:
		return true
	case int32:
		return true
	case int64:
		return true
	case string:
		_, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			return false
		}
		return true
	default:
		return false
	}
}

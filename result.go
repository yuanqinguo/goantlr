package goantlr

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

type Result struct {
	callIndex  int
	result     interface{}
	resultType ResultType
	err        Errno
}

func NewResult(result interface{}, callIndex int) *Result {
	resultType := ParseResultType(result)
	ret := &Result{
		result:     result,
		callIndex:  callIndex,
		resultType: resultType,
		err:        Pass,
	}
	return ret
}

func NewResultWithError(result interface{}, callIndex int, err error) *Result {
	ret := &Result{
		result:     result,
		callIndex:  callIndex,
		resultType: ResultError,
	}
	if err != nil {
		ret.err = Errno{
			Code:    -1,
			Message: err.Error(),
		}
	}
	return ret
}

func (r *Result) GetError() Errno {
	return r.err
}

func (r *Result) GetString() (string, error) {
	getType := reflect.TypeOf(r.result)
	switch getType.Kind() {
	case reflect.String:
		return r.result.(string), nil
	default:
		return fmt.Sprintf("%v", r.result), nil
	}
}

func (r *Result) GetBoolean() (bool, error) {
	if r.resultType != ResultBool {
		return false, fmt.Errorf("value type not bool")
	}
	if ret, ok := r.result.(bool); ok {
		return ret == true, nil
	}
	return false, nil
}

func (r *Result) GetInteger() (int64, error) {
	if r.resultType != ResultInteger {
		return 0, fmt.Errorf("value type not integer")
	}
	getType := reflect.TypeOf(r.result)
	switch getType.Kind() {
	case reflect.String:
		{
			num, err := strconv.Atoi(r.result.(string))
			if err != nil {
				return 0, err
			}
			return int64(num), nil
		}
	case reflect.Int:
		return int64(r.result.(int)), nil
	case reflect.Int32:
		return int64(r.result.(int32)), nil
	case reflect.Int64:
		return r.result.(int64), nil
	case reflect.Float32:
		return int64(r.result.(float32)), nil
	case reflect.Float64:
		return int64(r.result.(float64)), nil
	default:
		return 0, fmt.Errorf("value type error")
	}
}

func (r *Result) GetFloat() (float64, error) {
	if r.resultType != ResultFloat {
		return 0, fmt.Errorf("value type not float")
	}
	getType := reflect.TypeOf(r.result)
	switch getType.Kind() {
	case reflect.String:
		{
			float, err := strconv.ParseFloat(r.result.(string), 64)
			if err != nil {
				return 0, err
			}
			if math.IsNaN(float) {
				return 0, nil
			}
			return float, nil
		}
	case reflect.Int:
		return float64(r.result.(int)), nil
	case reflect.Int32:
		return float64(r.result.(int32)), nil
	case reflect.Int64:
		return float64(r.result.(int64)), nil
	case reflect.Float32:
		val, _ := r.result.(float32)
		if math.IsNaN(float64(val)) {
			return 0, nil
		}
		return float64(val), nil
	case reflect.Float64:
		val, _ := r.result.(float64)
		if math.IsNaN(val) {
			return 0, nil
		}
		return val, nil
	default:
		return 0, nil
	}
}

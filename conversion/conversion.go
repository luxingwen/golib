package conversion

import (
	"strconv"
)

func Int(val interface{}) int {
	switch val.(type) {
	case int:
		return val.(int)
	case int8:
		return int(val.(int8))
	case int16:
		return int(val.(int16))
	case int32:
		return int(val.(int32))
	case int64:
		return int(val.(int64))
	case uint:
		return int(val.(uint))
	case uint8:
		return int(val.(uint8))
	case uint16:
		return int(val.(uint16))
	case uint32:
		return int(val.(uint32))
	case uint64:
		return int(val.(uint64))
	case float32:
		return int(val.(float32))
	case float64:
		return int(val.(float64))
	case string:
		v, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			return 0
		}
		return int(v)
	default:
		return 0
	}
	return 0
}

func Int64(val interface{}) int64 {
	switch val.(type) {
	case int:
		return int64(val.(int))
	case int8:
		return int64(val.(int8))
	case int16:
		return int64(val.(int16))
	case int32:
		return int64(val.(int32))
	case int64:
		return val.(int64)
	case uint:
		return int64(val.(uint))
	case uint8:
		return int64(val.(uint8))
	case uint16:
		return int64(val.(uint16))
	case uint32:
		return int64(val.(uint32))
	case uint64:
		return int64(val.(uint64))
	case float32:
		return int64(val.(float32))
	case float64:
		return int64(val.(float64))
	case string:
		v, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			return 0
		}
		return v
	default:
		return 0
	}
	return 0
}

func Uint64(val interface{}) uint64 {
	switch val.(type) {
	case uint64:
		return val.(uint64)
	default:
		return uint64(Int64(val))
	}
	return 0
}

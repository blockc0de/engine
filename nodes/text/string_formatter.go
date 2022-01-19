package text

import (
	"fmt"
	"strconv"
	"strings"
)

type StringFormatter struct {
}

func (f StringFormatter) Format(format string, args interface{}) string {
	switch args := args.(type) {
	case []interface{}:
		return f.formatSlice(format, args...)
	case map[string]interface{}:
		return f.formatMap(format, args)
	default:
		return f.formatSlice(format, []interface{}{args})
	}
}

func (f StringFormatter) toString(v interface{}) (string, error) {
	switch v := v.(type) {
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case string:
		return v, nil
	case bool:
		return strconv.FormatBool(v), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}

func (f StringFormatter) formatSlice(format string, args ...interface{}) string {
	if args == nil {
		return format
	}
	for index, val := range args {
		arg := "{" + strconv.Itoa(index) + "}"
		if s, err := f.toString(val); err == nil {
			format = strings.Replace(format, arg, s, -1)
		}
	}
	return format
}

func (f StringFormatter) formatMap(format string, args map[string]interface{}) string {
	if args == nil {
		return format
	}
	for key, val := range args {
		arg := "{" + key + "}"
		if s, err := f.toString(val); err == nil {
			format = strings.Replace(format, arg, s, -1)
		}
	}
	return format
}

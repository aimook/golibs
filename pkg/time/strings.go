package t

import (
	"fmt"
	"strconv"
)

func ToString(source interface{}) string {
	switch source.(type) {
	case string:
		return source.(string)
	case bool:
		return strconv.FormatBool(source.(bool))
	case int:
		return strconv.Itoa(source.(int))
	case float32:
		return fmt.Sprintf("%f", source.(float32))
	case float64:
		return fmt.Sprintf("%f", source.(float64))
	default:
		return fmt.Sprintf("%v", source)
	}
}

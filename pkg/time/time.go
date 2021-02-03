package t

import (
	"strconv"
	"time"
)

type Time int64

const (
	timeFormat = "2006-01-02 15:04:05"
	zone       = "Asia/Shanghai"
)

//Scan 数据库->应用程序自定义时间类型转换
func (r *Time) Scan(v interface{}) error {
	vv, ok := v.(int64)
	if ok {
		*r = Time(vv)
	}
	return nil
}

//UnmarshalJSON 字符串JSON->对象自定义时间类型转换
func (r *Time) UnmarshalJSON(bytes []byte) error {
	val, _ := strconv.ParseInt(string(bytes), 10, 64)
	*r = Time(val)
	return nil
}

//MarshalJSON 对象自定义时间类型 -> JSON对象转换
func (r Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Unix(int64(r), 0).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

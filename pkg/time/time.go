package utils

import (
	"database/sql/driver"
	"time"
)

type ParkTime time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
	zone       = "Asia/Shanghai"
)

func (t *ParkTime) Scan(v interface{}) error {
	value, ok := v.(int64)
	if ok {
		*t = ParkTime(time.Unix(value, 0))
	}
	return nil
}

func (t ParkTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *ParkTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = ParkTime(now)
	return
}

func (t ParkTime) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t ParkTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

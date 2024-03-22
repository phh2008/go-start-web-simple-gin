package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	if data == nil || string(data) == "null" {
		return
	}
	if len(data) == 2 {
		*t = DateTime(time.Time{})
		return
	}
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = DateTime(now)
	return
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	value := time.Time(t)
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = value.AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t DateTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

func (t *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to time.Time", v)
}

func (t DateTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

func (t *DateTime) ToTime() time.Time {
	return time.Time(*t)
}

package util

import (
	"fmt"
	"time"
)

type JsonTime struct {
	time.Time
}

func (jsonTime *JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", jsonTime.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (jsonTime *JsonTime) UnmarshalJSON(b []byte) error {
	parsed, err := time.Parse("2006-01-02 15:04:05", string(b))
	if err != nil {
		return err
	}
	jsonTime.Time = parsed
	return nil
}

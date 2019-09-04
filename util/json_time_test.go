package util

import (
	"fmt"
	"testing"
	"time"
)

func TestJsonTime_MarshalJSON(t *testing.T) {
	jsonTime := JsonTime{time.Now()}
	marshal, _ := jsonTime.MarshalJSON()
	fmt.Println(string(marshal))
}

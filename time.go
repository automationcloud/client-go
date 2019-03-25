package client

import (
	"encoding/json"
	"time"
)

type jsTime struct {
	time.Time
}

func (t *jsTime) UnmarshalJSON(d []byte) (err error) {
	var ms int64
	err = json.Unmarshal(d, &ms)
	if err != nil {
		return
	}

	t.Time = time.Unix(0, ms*int64(time.Millisecond))
	return
}

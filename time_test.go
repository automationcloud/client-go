package client

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestTimeUnmarshaler(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		var tm jsTime
		if err := json.Unmarshal([]byte("0"), &tm); err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		time := tm.UTC().Format(time.RFC3339Nano)
		expected := "1970-01-01T00:00:00Z"
		if time != expected {
			t.Errorf("Expected time to be %v, got: %v", expected, time)
		}

	})

	t.Run("unhappy case", func(t *testing.T) {
		var tm jsTime
		err := json.Unmarshal([]byte("{}"), &tm)
		if err == nil {
			t.Error("Expected error, got success")
			t.FailNow()
		}

		expected := "cannot unmarshal object into Go value of type int64"
		if !strings.Contains(err.Error(), expected) {
			t.Errorf("Expected error to be %v, got: %v", expected, err)
		}
	})
}

package gorunoutside

import (
  "testing"
  "json"
  "fmt"
)

func TestJSONMarshalling(t *testing.T) {
  activity := Activity{"Running", "Sat, 1 Jan 2011 00:00:00", 3.4, nil, 1000, 355.0}
  outputJSON, _ := json.Marshal(activity)
  fmt.Printf("JSON Output: %s", outputJSON)
}
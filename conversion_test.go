package main

import (
  "testing"
  "github.com/bmizerany/assert"
  "os"
  "xml"
  "math"
)

func TestConvertExtendedData(t *testing.T) {
  distances := ConvertExtendedData("0.0, 0.0182, 0.0456, 0.0703")
  assert.Equal(t, 4, len(distances))
  assert.Equal(t, 0, distances[0].Timestamp)
  assert.Equal(t, 0.0, distances[0].Distance)
  assert.Equal(t, 10, distances[1].Timestamp)
  assert.Equal(t, 18.2, distances[1].Distance)
  assert.Equal(t, 20, distances[2].Timestamp)
  assert.Equal(t, 45.6, distances[2].Distance)
}

func TestConvertNikePlusToRunKeeper(t *testing.T) {
  file, err := os.Open("fixtures/individual_run.xml")
  if err == nil {
    ps := new(PlusService)
    xml.Unmarshal(file, ps)
    file.Close()
    
    activity := ConvertNikePlusToRunKeeper(ps)
    
    assert.Equal(t, "Running", activity.Type)
    assert.Equal(t, "Sun, 09 Dec 2011 07:11:45 -0500", activity.StartTime)
    assert.Equal(t, 3951.9, activity.TotalDistance)
    assert.Equal(t, 160, len(activity.Distance))
    assert.Equal(t, 18.0, math.Floor(activity.Distance[1].Distance))
    assert.Equal(t, 3925.0, math.Floor(activity.Distance[159].Distance))
  } else {
    t.Error(err)
  }
}
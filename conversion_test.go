package main

import (
  "testing"
  "github.com/bmizerany/assert"
)

func TestConvertExtendedData(t *testing.T) {
  distances := *ConvertExtendedData("0.0, 0.0182, 0.0456, 0.0703")
  assert.Equal(t, 4, len(distances))
  assert.Equal(t, 0, distances[0].Timestamp)
  assert.Equal(t, 0.0, distances[0].Distance)
  assert.Equal(t, 10, distances[1].Timestamp)
  assert.Equal(t, 18.2, distances[1].Distance)
  assert.Equal(t, 20, distances[2].Timestamp)
  assert.Equal(t, 45.6, distances[2].Distance)
}
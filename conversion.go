package main

import (
  "strings"
  "strconv"
)

func ConvertExtendedData(extendedData string) *[]Distance {
  kmStrings := strings.Split(extendedData, ", ")
  distances := make([]Distance, len(kmStrings))
  for i := 0; i < len(kmStrings); i++ {
    kmDistance, _ := strconv.Atof64(kmStrings[i])
    distances[i] = Distance{i * 10, kmDistance * 1000}
  }
  return &distances
}
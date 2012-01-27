package main

import (
  "strings"
  "strconv"
  "time"
  "fmt"
)

func ConvertExtendedData(extendedData string) []Distance {
  kmStrings := strings.Split(extendedData, ", ")
  distances := make([]Distance, len(kmStrings))
  for i := 0; i < len(kmStrings); i++ {
    kmDistance, _ := strconv.Atof64(kmStrings[i])
    distances[i] = Distance{i * 10, kmDistance * 1000}
  }
  return distances
}

func ConvertNikePlusToRunKeeper(plusService *PlusService) *Activity {
  activity := new(Activity)
  activity.Type = "Running"
  t := plusService.SportsData.ParsedStartTime()
  if t != nil {
    activity.StartTime = t.Format(time.RFC1123)
  } else {
    fmt.Print("Couldn't find a start time for run")
  }
  activity.TotalDistance = plusService.SportsData.RunSummary.Distance * 1000
  activity.Distance = ConvertExtendedData(plusService.SportsData.ExtendedDataList.ExtendedData)
  activity.Duration = plusService.SportsData.RunSummary.Duration
  activity.TotalCalories = plusService.SportsData.RunSummary.Calories
  return activity
}
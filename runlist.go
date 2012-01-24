package main

type Run struct {
  Id int `xml:"attr"`
  StartTime string
  Distance float64
  Duration int
}

type RunList struct {
  Run []Run
}

type Stats struct {
  Distance float64
  Calories float64
  Duration int
}

func (s Stats) DistanceInMiles() float64 {
  return s.Distance * 0.621371192
}

type RunListSummary struct {
  Stats
  Runs int
  RunDuration int
}

type UserInfo struct {
  Weight float64
  Device string
}

type RunSummary struct {
  Stats
  EquipmentType string
  HasGpsData bool
}

type ExtendedDataList struct {
  ExtendedData string
}

type SportsData struct {
  UserInfo UserInfo
  RunSummary RunSummary
  ExtendedDataList ExtendedDataList
}

type PlusService struct {
  RunList RunList
  RunListSummary RunListSummary
  SportsData SportsData
  Status string
}
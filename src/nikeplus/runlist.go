package nikeplus

type Run struct {
  Id int `xml:"attr"`
  StartTime string
  Distance float64
  Duration int
}

type RunList struct {
  Run []Run
}

type RunListSummary struct {
  Runs int
  Distance float64
  RunDuration int
  Calories float64
  Duration int
}

func (rls RunListSummary) DistanceInMiles() float64 {
  return rls.Distance * 0.621371192
}

type PlusService struct {
  RunList RunList
  RunListSummary RunListSummary
  Status string
}
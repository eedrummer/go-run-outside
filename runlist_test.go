package main

import (
  "testing"
  "os"
  "xml"
  "github.com/bmizerany/assert"
)

func TestLoadAllRuns(t *testing.T) {
  file, err := os.Open("fixtures/example.xml")
  if err == nil {
    ps := new(PlusService)
    xml.Unmarshal(file, ps)
    file.Close()
    
    firstRun := &ps.RunList.Run[0]
    assert.Equal(t, 1484663528, firstRun.Id)
  } else {
    t.Error(err)
  }
}

func TestLoadIndividualRun(t *testing.T) {
  file, err := os.Open("fixtures/individual_run.xml")
  if err == nil {
    ps := new(PlusService)
    xml.Unmarshal(file, ps)
    file.Close()
    
    runSummary := &ps.SportsData.RunSummary
    assert.Equal(t, 3.9519, runSummary.Distance)
  } else {
    t.Error(err)
  }
}

func TestDistanceInMiles(t *testing.T) {
  rls := new(RunListSummary)
  rls.Distance = 15
  assert.Equal(t, 9.32056788, rls.DistanceInMiles())
}

func TestSportsDataTimeConversion(t *testing.T) {
  sd := new(SportsData)
  sd.StartTime = "2011-12-09T07:11:45-05:00"
  startTime := sd.ParsedStartTime()
  assert.Equal(t, 12, startTime.Month)
  assert.Equal(t, 9, startTime.Day)
}
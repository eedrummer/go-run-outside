package nikeplus

import (
  "testing"
  "os"
  "xml"
)

func TestLoadXml(t *testing.T) {
  file, err := os.Open("../../fixtures/example.xml")
  if err == nil {
    ps := new(PlusService)
    xml.Unmarshal(file, ps)
    file.Close()
    
    firstRun := &ps.RunList.Run[0]
    expectedId := 1484663528
    if firstRun.Id != expectedId {
      t.Errorf("Expected %d but got %d", expectedId, firstRun.Id)
    }
  } else {
    t.Error(err)
  }
}

func TestDistanceInMiles(t *testing.T) {
  rls := new(RunListSummary)
  rls.Distance = 15
  expectedDistance := 9.32056788
  if rls.DistanceInMiles() != expectedDistance {
    t.Errorf("Expected %d but got %d", expectedDistance, rls.DistanceInMiles())
  }
}
package nikeplus

import (
  "testing"
  "os"
  "xml"
)

func TestLoadXml(t *testing.T) {
  file, err := os.Open("../../fixtures/example.xml")
  if err == nil {
    rl := RunList{Run: nil}
    xml.Unmarshal(file, &rl)
    file.Close()
    
    firstRun := &rl.Run[0]
    expectedId := 1484663528
    if firstRun.Id != expectedId {
      t.Errorf("Expected %d but got %d", expectedId, firstRun.Id)
    }
  } else {
    t.Error(err)
  }
}
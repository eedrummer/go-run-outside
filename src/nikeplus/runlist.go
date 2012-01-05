package nikeplus

import (
  "xml"
)

type Run struct {
  Id int `xml:"attr"`
  StartTime string
  Distance float32
  Duration int
}

type RunList struct {
  XMLName xml.Name `xml:"runList"`
  Run []Run
}
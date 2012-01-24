package main

type Distance struct {
  Timestamp int `json:"timestamp"`
  Distance float64 `json:"distance"`
}

type Activity struct {
  Type string `json:"type"`
  StartTime string `json:"start_time"`
  TotalDistance float64 `json:"total_distance"`
  Distance []Distance `json:"distance"`
  Duration int `json:"duration"`
  TotalCalories float64 `json:"total_calories"`
}
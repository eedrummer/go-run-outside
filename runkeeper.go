package main

import (
  "http"
  "json"
  "fmt"
  "strings"
)

func PostActivity(activity *Activity, bearerToken string, client *http.Client) {
  outputJSON, _ := json.Marshal(activity)
  request, _ := http.NewRequest("POST",  "https://api.runkeeper.com/fitnessActivities", strings.NewReader(string(outputJSON)))
  request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
  request.Header.Add("Content-Type", "application/vnd.com.runkeeper.NewFitnessActivity+json")
  response, _ := client.Do(request)
  if response.StatusCode == 201 {
    fmt.Print("Activity POSTed")
  } else {
    fmt.Printf("Activity post failed with status code %d and message %s", response.StatusCode, response.Body)
  }
}
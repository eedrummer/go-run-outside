package main

import (
  "http"
  "xml"
  "fmt"
)

func GetRunList(client *http.Client, userId string) *PlusService {
  url := fmt.Sprintf("http://nikerunning.nike.com/nikeplus/v1/services/widget/get_public_run_list.jsp?userID=%s", userId)
  return fetchRunList(client, url)
}

func GetIndividualRun(client *http.Client, userId string, runId int) *PlusService {
  url := fmt.Sprintf("http://nikerunning.nike.com/nikeplus/v2/services/app/get_public_run.jsp?id=%d&userID=%s", runId, userId)
  return fetchRunList(client, url)
}

func fetchRunList(client *http.Client, url string) *PlusService {
  response, _ := client.Get(url)
  ps := new(PlusService)
  xml.Unmarshal(response.Body, ps)
  return ps
}
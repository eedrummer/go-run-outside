package main

import (
  "flag"
  "fmt"
  "http"
  "time"
)

func main() {
  uid := GetUid()
  
  client := new(http.Client)
  runChannel := make(chan *PlusService, 5)
  countChannel := make(chan int)
  go func() {
    bigList := GetRunList(client, uid)
    numberOfRuns := len(bigList.RunList.Run)
    countChannel <- numberOfRuns
    for i := 0; i < numberOfRuns; i++ {
      runChannel <- GetIndividualRun(client, uid, bigList.RunList.Run[i].Id)
    }
  }()
  
  bearerToken := CheckForBearerToken()
  if bearerToken == "" {
    LaunchOAuth()
    for bearerToken == "" {
      time.Sleep(500000000)
      bearerToken = CheckForBearerToken()
    }
  }
  
  runCount := <- countChannel
  for i := 0; i < runCount; i++ {
    ps := <- runChannel
    activity :=  ConvertNikePlusToRunKeeper(ps)
    PostActivity(activity, bearerToken, client)
  }
}

func GetUid() string {
  var uid string
  flag.StringVar(&uid, "u", "", "The Nike+ user id")
  flag.Parse()
  if uid == "" {
    panic("No user id provided")
  }
  fmt.Printf("Nike+ user %s was passed in.\n", uid)
  return uid
}

func LaunchOAuth() {
  fmt.Print("No bearer token found, going through the OAuth process.\n")
  http.HandleFunc("/code", OAuthCallbackServerHelloServer)
  go http.ListenAndServe(":4444", nil)
  OpenBrowser()
}
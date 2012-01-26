package main

import (
  "exec"
  "http"
  "os"
  "fmt"
  "json"
  "io/ioutil"
)

const (
  RedirectUri = "http://localhost:4444/code"
  ClientId = "ef74ff22481040419162523f2ab0d65c"
  ClientSecret = "3fb54cc1a94142bc87a824d1ccb59911"
)

func OpenBrowser() {
  url := fmt.Sprintf("https://runkeeper.com/apps/authorize?client_id=%s&response_type=code&redirect_uri=%s", ClientId, RedirectUri)
  command := exec.Command("open", url)
  command.Run()
}

func OAuthCallbackServerHelloServer(w http.ResponseWriter, req *http.Request) {
  code := req.URL.Query().Get("code")
  go ObtainBearerToken(code)
  w.Write([]uint8("Called Back!\n"))
}

func ObtainBearerToken(code string) {
  tokenUrl := "https://runkeeper.com/apps/token"
  formData := make(map[string][]string)
  formData["grant_type"] = []string{"authorization_code"}
  formData["code"] = []string{code}
  formData["client_id"] = []string{ClientId}
  formData["client_secret"] = []string{ClientSecret}
  formData["redirect_uri"] = []string{RedirectUri}
  client := new(http.Client)
  response, err := client.PostForm(tokenUrl, formData)
  responseJson := make(map[string]string)
  if err == nil {
    responseBody, _ := ioutil.ReadAll(response.Body)
    json.Unmarshal(responseBody, &responseJson)
    file, _ := os.OpenFile(".bearer_token", os.O_TRUNC | os.O_WRONLY, 644)
    file.WriteString(responseJson["access_token"])
    file.Close()
  } else {
    fmt.Print(err)
  }
}
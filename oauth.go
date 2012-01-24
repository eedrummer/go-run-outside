package main

import (
  "exec"
)

func OpenBrowser() {
  url := "https://runkeeper.com/apps/authorize?client_id=ef74ff22481040419162523f2ab0d65c&response_type=code&redirect_uri=http://localhost:4444/code"
  command := exec.Command("open", url)
  command.Run()
}
package main

import (
  "testing"
  "os"
  "github.com/bmizerany/assert"
)

func TestCheckForBearerTokenWithoutFile(t *testing.T) {
  bt := CheckForBearerToken()
  if bt != "" {
    t.Error("Bearer token should be nil")
  }
}

func TestCheckForBearerTokenWithFile(t *testing.T) {
  file, _ := os.Create(".bearer_token")
  file.WriteString("123456\n")
  file.Close()
  bt := CheckForBearerToken()
  os.Remove(".bearer_token")
  assert.Equal(t, "123456", bt)
}

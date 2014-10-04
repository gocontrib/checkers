package gocheckers

import (
  "os"
  . "gopkg.in/check.v1"
)

type existsChecker struct {}

// Checker impl
func (c *existsChecker) Info() *CheckerInfo {
  return &CheckerInfo{Name: "exists", Params: []string{"value","type"}}
}

func (c *existsChecker) Check(params []interface{}, names []string) (result bool, error string) {
  path, ok := params[0].(string)
  if !ok {
    error = "expected string argument"
    return
  }
  kind, ok2 := params[1].(string)
  if (!ok2) {
    kind = "file"
  }
  isDir := kind == "dir"
  isFile := kind == "file"
  s, err := os.Stat(path)
  result = err != nil && s.IsDir() == isDir && s.IsDir() != isFile
  return
}

var Exists Checker = &existsChecker{}

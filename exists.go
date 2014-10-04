package gocheckers

import (
  "os"
  . "gopkg.in/check.v1"
)

// The Exists checker verifies that the obtained path to dir/file exists
//
// Examples:
//
// 		c.Check(path, Exists) - checks file system item exists
// 		c.Check(path, Exists, "dir") - checks file system item exists and it is directory
// 		c.Check(path, Exists, "file") - checks file system item exists and it is file
//
var Exists Checker = &existsChecker{}

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

  st, err := os.Stat(path)
  if (err != nil) {
  	result = false
  	return
  }

  if len(params) == 2 {
  	kind, ok2 := params[1].(string)
  	if (!ok2) {
    	result = false
    	return
  	}
  	if st.IsDir() {
  		result = kind == "dir"
  		return
  	}
  	result = kind == "file"
  }

  return
}

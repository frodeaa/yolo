package main

import "github.com/frodeaa/yolo/cmd"

var (
	VERSION = "0.0.1"
)

func main() {
	cmd.Execute(VERSION)
}

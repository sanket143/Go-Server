package main

import "github.com/sanket143/lisa/src/system/app"

func main() {
	s := app.NewServer()

	s.Init()
	s.Start()
}

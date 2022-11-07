package main

import (
	"dyslexia/conf"
	"dyslexia/server"
)

func main() {
	conf.Load()
	server.Start()
}

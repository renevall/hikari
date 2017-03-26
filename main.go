package main

import "bitbucket.org/reneval/hikari/service"

//TODO: makes this file a cli

func main() {
	hs := service.HikariService{}
	hs.Run()
}

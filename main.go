package main

import (
	"anonymousroom/gata"
	"github.com/name5566/leaf"
	"github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
)

func main() {
	log.Debug("enter main")
	defer log.Debug("level main")

	conf.LogLevel = "debug"

	g := gata.NewGateModule("", ":8889")
	leaf.Run(g)
}

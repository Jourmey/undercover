package main

import (
	"github.com/name5566/leaf"
	"github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

//var Chanel = chanrpc.NewServer(1000)

func main() {
	log.Debug("enter main")
	defer log.Debug("level main")

	conf.LogLevel = "debug"

	//Chanel.Register(reflect.TypeOf(&Hello{}), HelloHandle)

	//m := NewGameModule()
	g := NewGateModule()

	Processor.Register(&Hello{})
	//Processor.SetRouter(&Hello{}, Chanel)
	Processor.SetHandler(&Hello{}, HelloHandle)

	leaf.Run(
		//m,
		g)
}

func HelloHandle(args []interface{}) {
	log.Release("enter HelloHandle %+v", args)
}

type Hello struct {
	Name string
}

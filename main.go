package main

import (
	"flag"
	"github.com/ukinhappy/webhook/config"
	"github.com/ukinhappy/webhook/hook"
	"github.com/ukinhappy/webhook/logger"
	"github.com/ukinhappy/webhook/server"
	"log"
)

var (
	cfpath = flag.String("config", "", "--config=")
)

func main() {
	flag.Parse()
	if *cfpath == "" {
		log.Fatal("config can not nil")
	}
	config.Init(*cfpath)
	logger.Init()
	hook.LoadAllWebHook()
	server.Server()
}

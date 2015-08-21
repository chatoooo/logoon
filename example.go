package main

import (
	log "github.com/chatoooo/logoon/log"
	logSource "github.com/chatoooo/logoon/source"
	"time"
)

func main() {
	dispatcher, err := log.CreateLoggingFromFile("logoon.json")
	if err != nil {
		panic(err)
	}

	accessLog := logSource.MakeTaggedLogSource(dispatcher, []string{"access"})
	accessLog.Info("Access I") //this should be visible in ./access.log
	accessLog.Error("Access E") //this should be visible in ./access.log

	consoleLog := logSource.MakeTaggedLogSource(dispatcher, []string{"console"})
	consoleLog.Info("Console Info") //this should be visible in console
	consoleLog.Error("Console Error") //this should be visible in console

	errorLog := logSource.MakeSimpleLogSource(dispatcher)
	errorLog.Info("Error Info") //this be discarded
	errorLog.Error("Error Error") //this should be visible in ./error.log

	time.Sleep(100)
}

package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/cihub/seelog"
	"github.com/igoboy/httpcmd"
	"github.com/igoboy/httpcmd/log"
)

func help() {
	fmt.Printf(
		"%s ip:port\n",
		os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	defer seelog.Flush()
	logFile := "./httpcmd.log"
	if log.InitLog("debug", logFile) < 0 {
		fmt.Println("Init log file error ", logFile)
		return
	}

	addr := os.Args[1]
	httpcmd.StartHttpServer(addr)

	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)

	signal.Notify(signalChan,
		syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	signal.Stop(signalChan)
	seelog.Info("See you next time at httpcmd !")
}

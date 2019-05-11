package httpcmd

import (
	"net/http"

	"github.com/cihub/seelog"
)

var HttpListenAddr string

func StartHttpServer(listenAddr string) int {
	if listenAddr == "" {
		return 0
	}

	go startHttpServer(listenAddr)
	return 0
}

func startHttpServer(listenAddr string) int {
	HttpListenAddr = listenAddr
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		seelog.Error("ListenAndServe: ", err)
		return -1
	}

	return 0
}

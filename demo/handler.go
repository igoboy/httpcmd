package main

import (
	"net/http"
)

func init() {
	http.Handle("/demo/command", MakeCommandHandler())
}

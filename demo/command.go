package main

import (
	"net/http"

	"github.com/cihub/seelog"
	"github.com/igoboy/httpcmd/base"
)

type CommandHandler struct {
	base.BasicHandler
}

func MakeCommandHandler() http.Handler {
	h := &CommandHandler{}
	h.HandlerInit(h)
	return h
}

func (this *CommandHandler) ProcessAdd(context *base.HttpContext) error {
	seelog.Infof("implement command add")
	return nil
}

func (this *CommandHandler) ProcessSend(context *base.HttpContext) error {
	seelog.Infof("implement command send")
	return nil
}

func (this *CommandHandler) ProcessExtend(context *base.HttpContext) error {
	if context.GetCmd() == "send" {
		return this.ProcessSend(context)
	}

	return this.BasicHandler.ProcessExtend(context)
}

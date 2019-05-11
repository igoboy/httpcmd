package base

import (
	"fmt"
	"net/http"

	"github.com/cihub/seelog"
)

type Handler interface {
	HandlerInit(h Handler)
	ProcessAdd(context *HttpContext) error
	ProcessModify(context *HttpContext) error
	ProcessUpdate(context *HttpContext) error
	ProcessDelete(context *HttpContext) error
	ProcessExtend(context *HttpContext) error
	ProcessRequest(context *HttpContext) error
	ProcessResponse(context *HttpContext) error
}

type BasicHandler struct {
	BaseHandler
	vptr Handler
}

func (this *BasicHandler) HandlerInit(h Handler) {
	this.vptr = h
}

func (this *BasicHandler) ServeHTTP(
	rsp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	context := &HttpContext{
		HttpReq: req,
		HttpRsp: rsp,
		ReqData: &CommandRequest{},
		RspData: &CommandResponse{},
	}
	this.BaseHandler.ServeHTTP(rsp, req)

	if this.vptr == nil {
		this.ProcessResponse(
			MakeFailure(context,
				fmt.Errorf("not support handler")),
		)
		return
	}

	if err := this.ProcessRequest(context); err != nil {
		this.ProcessResponse(MakeFailure(context, err))
		return
	}

	this.ProcessResponse(context)
}

func (this *BasicHandler) ProcessAdd(context *HttpContext) error {
	return this.processCmdImpl(context)
}

func (this *BasicHandler) ProcessModify(context *HttpContext) error {
	return this.processCmdImpl(context)
}

func (this *BasicHandler) ProcessUpdate(context *HttpContext) error {
	return this.processCmdImpl(context)
}

func (this *BasicHandler) ProcessDelete(context *HttpContext) error {
	return this.processCmdImpl(context)
}

func (this *BasicHandler) ProcessExtend(context *HttpContext) error {
	return this.processCmdImpl(context)
}

func (this *BasicHandler) processCmdImpl(context *HttpContext) error {
	return fmt.Errorf(
		"not support %s cmd %s", context.GetCmd(), context.ID())
}

func (this *BasicHandler) ProcessRequest(context *HttpContext) error {
	if err := this.recvData(context.HttpReq,
		context.ReqData); err != nil {
		seelog.Errorf("recv data error %s %s", err, context.ID())
		MakeFailure(context, err)
		return err
	}

	MakeSuccess(context)
	if err := context.Init(); err != nil {
		MakeFailure(context, err)
		return err
	}

	if context.GetCmd() == "add" {
		return this.vptr.ProcessAdd(context)
	}

	if context.GetCmd() == "modify" {
		return this.vptr.ProcessModify(context)
	}

	if context.GetCmd() == "update" {
		return this.vptr.ProcessUpdate(context)
	}

	if context.GetCmd() == "delete" {
		return this.vptr.ProcessDelete(context)
	}

	return this.vptr.ProcessExtend(context)
}

func (this *BasicHandler) ProcessResponse(context *HttpContext) error {
	if err := this.sendData(
		context.HttpRsp, context.RspData); err != nil {
		seelog.Errorf("send data error %s %s", err, context.ID())
		return err
	}

	return nil
}

package base

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golibs/uuid"
)

type HttpContext struct {
	HttpReq *http.Request
	HttpRsp http.ResponseWriter

	ReqData *CommandRequest
	RspData *CommandResponse
}

func (this *HttpContext) ID() string {
	return this.ReqData.Id
}

func UUID() string {
	return strings.Replace(
		uuid.Rand().Hex(), "-", "", -1)
}

func (this *HttpContext) Init() error {
	if this.ReqData.Id == "" {
		return fmt.Errorf(
			"reqeust id is empty")
	}

	this.RspData.Id = this.ReqData.Id
	if this.ReqData.Cmd == "" {
		return fmt.Errorf(
			"request cmd is empty")
	}

	return nil
}

func (this *HttpContext) GetCmd() string {
	return this.ReqData.Cmd
}

func MakeFailure(
	context *HttpContext, err error) *HttpContext {
	return MakeFailureResponse(context, -1, err, nil)
}

func MakeFailure2(
	context *HttpContext, code int, err error) *HttpContext {
	return MakeFailureResponse(context, code, err, nil)
}

func MakeFailureResponse(context *HttpContext,
	code int, err error, data interface{}) *HttpContext {
	if context.ReqData == nil ||
		context.RspData == nil {
		return context
	}

	id := context.ID()
	if id == "" {
		id = UUID()
	}

	context.RspData.Id = id
	if context.RspData.Code < -1 {
		return context
	}

	context.RspData.Code = code
	context.RspData.Data = data
	context.RspData.Msg = fmt.Sprintf("%s", err)
	return context
}

func MakeSuccess(context *HttpContext) *HttpContext {
	return MakeSuccessResponse(context, nil)
}

func MakeSuccessResponse(
	context *HttpContext, data interface{}) *HttpContext {
	if context.ReqData == nil ||
		context.RspData == nil {
		return context
	}

	context.RspData.Code = 0
	context.RspData.Data = data
	context.RspData.Msg = "success"
	context.RspData.Id = context.ID()
	return context
}

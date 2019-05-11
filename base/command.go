package base

import "encoding/json"

type CommandRequest struct {
	Id   string          `json:"id"`
	Cmd  string          `json:"command"`
	Data json.RawMessage `json:"data,omitempty"`
}

type CommandResponse struct {
	Id   string      `json:"id"`
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

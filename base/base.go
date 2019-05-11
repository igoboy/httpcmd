package base

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/cihub/seelog"
)

type BaseHandler struct{}

var defaultHandler BaseHandler

func ReadData(r io.Reader, d interface{}) error {
	return defaultHandler.ReadData(r, d)
}

func WriteData(w io.Writer, d interface{}) error {
	return defaultHandler.WriteData(w, d)
}

func (this *BaseHandler) ReadData(r io.Reader, d interface{}) error {
	body, err := ioutil.ReadAll(r)
	if nil != err {
		return err
	}

	seelog.Infof("recv body %s", body)
	err = json.Unmarshal(body, d)
	if nil != err {
		return err
	}

	return nil
}

func (this *BaseHandler) WriteData(w io.Writer, d interface{}) error {
	body, err := json.Marshal(d)
	if nil != err {
		return err
	}

	seelog.Infof("send body %s", string(body))
	_, err = this.Write(w, body)
	if err != nil {
		return err
	}

	return nil
}

func (this *BaseHandler) Write(w io.Writer, data []byte) (int, error) {
	n := 0
	length := len(data)
	for length > n {
		ret, err := w.Write(data[n:])
		if err != nil {
			return n, err
		}

		n += ret
	}

	return n, nil
}

func (this *BaseHandler) recvData(req *http.Request, d interface{}) error {
	return this.ReadData(req.Body, d)
}

func (this *BaseHandler) sendData(rsp http.ResponseWriter, d interface{}) error {
	return this.WriteData(rsp, d)
}

func (this *BaseHandler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	rsp.Header().Set("Content-Type", "application/json")
}

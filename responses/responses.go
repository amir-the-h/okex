package responses

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
)

type (
	RestBasic struct {
		Code int    `json:"code,string"`
		Msg  string `json:"msg,omitempty"`
	}

	WsBasicEvent struct {
		Event string         `json:"event"`
		Code  int            `json:"code,omitempty,string"`
		Msg   string         `json:"msg,omitempty"`
		Op    okex.Operation `json:"op,omitempty"`
		Arg   *WsArgument    `json:"arg,omitempty"`
		Args  []*WsArgument  `json:"args,omitempty"`
		Data  []*WsArgument  `json:"data,omitempty"`
	}
	WsArgument struct {
		arg map[string]interface{}
	}

	WsError struct {
		Event string `json:"event"`
		Code  string `json:"code"`
		Msg   string `json:"msg"`
	}

	WsSubscribe struct {
		Event string      `json:"event"`
		Arg   *WsArgument `json:"arg"`
	}
	WsUnsubscribe struct {
		Event string      `json:"event"`
		Arg   *WsArgument `json:"arg"`
	}
)

func (a *WsArgument) Get(k string) (interface{}, bool) {
	v, ok := a.arg[k]
	return v, ok
}

func (a *WsArgument) UnmarshalJSON(buf []byte) error {
	a.arg = make(map[string]interface{})
	return json.Unmarshal(buf, &a.arg)
}
